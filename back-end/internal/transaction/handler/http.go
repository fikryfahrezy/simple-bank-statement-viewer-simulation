package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
)

type TransactionHandler struct {
	transactionService service.TransactionService
	log                *slog.Logger
}

func NewTransactionHandler(log *slog.Logger, transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		log:                log,
	}
}

// translateServiceError converts service errors to appropriate HTTP responses
func (h *TransactionHandler) translateServiceError(w http.ResponseWriter, err error, defaultMessage string) {
	if errors.Is(err, repository.ErrTransactionsTableNotFound) {
		http_server.InternalServerErrorResponse(w, "Transaction table uninitialize", err)
	}

	// Log unexpected errors
	h.log.Error("Service error",
		slog.String("error", err.Error()),
		slog.String("operation", defaultMessage),
	)
	http_server.InternalServerErrorResponse(w, defaultMessage, err)
}

// Setup HTTP API houtes for transactions
func (h *TransactionHandler) SetupRoutes(server *http_server.Server) {
	server.HandleFunc("POST /upload", h.Upload)
	server.HandleFunc("GET /balance", h.Balance)
	server.HandleFunc("GET /issues", h.Issues)
}
