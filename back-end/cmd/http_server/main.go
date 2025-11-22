// @title Go Simple API
// @description A simple API writte in Go
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @schemes http https
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/config"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"

	_ "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/docs"
	healthHandler "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/health/handler"
	server "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"

	transactionHandler "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	transactionRepository "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	transactionService "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildTime = "unknown"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Logger)

	log.Info("Starting application",
		slog.String("version", version),
		slog.String("commit", commit),
		slog.String("build_time", buildTime),
		slog.String("server_host", cfg.Server.Host),
		slog.Int("server_port", cfg.Server.Port),
	)

	db, err := database.NewDB(map[any]map[any]any{
		"transactions": {},
	})
	if err != nil {
		log.Error("Failed to connect to database",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	// Create server configuration
	serverConfig := server.Config{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
	}

	// Initialize feature dependencies
	transactionRepo := transactionRepository.NewTransactionRepository(log, db)
	transactionService := transactionService.NewTransactionService(log, transactionRepo)
	transactionHandler := transactionHandler.NewTransactionHandler(log, transactionService)

	// Initialize health handler
	healthHandler := healthHandler.NewHealthHandler(db, version, commit, buildTime)

	// Create and initialize server
	srv := server.New(serverConfig)

	routeHandlers := []server.RouteHandler{
		healthHandler,
		transactionHandler,
	}

	// Start server in goroutine
	go func() {
		if err := srv.Start(routeHandlers); err != nil {
			log.Error("Server error",
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down gracefully...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop server
	if err := srv.Stop(ctx); err != nil {
		log.Error("Failed to shutdown server gracefully",
			slog.String("error", err.Error()),
		)
	}

	log.Info("Application shutdown complete")
}
