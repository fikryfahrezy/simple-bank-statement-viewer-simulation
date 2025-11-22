package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_GetIssues_Success(t *testing.T) {
	db, err := database.NewDB(map[any]map[any]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	transactions := db.Table["transactions"]

	firstTransaction := model.Transaction{
		ID:        uuid.New(),
		Name:      "Transaction 1",
		Email:     "transaction1@example.com",
		Password:  "hashedpassword1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	transactions[firstTransaction.ID.String()] = firstTransaction

	secondTransaction := model.Transaction{
		ID:        uuid.New(),
		Name:      "Transaction 2",
		Email:     "transaction2@example.com",
		Password:  "hashedpassword2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	transactions[secondTransaction.ID.String()] = secondTransaction

	paginationReq := service.GetIssuesRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}

	result, totalCount, err := transactionService.GetIssues(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(2), totalCount)

	// Verify first transaction
	assert.Equal(t, firstTransaction.ID, result[0].ID)
	assert.Equal(t, firstTransaction.Name, result[0].Name)
	assert.Equal(t, firstTransaction.Email, result[0].Email)
}

func TestTransactionService_GetIssues_WithCustomPagination(t *testing.T) {
	db, err := database.NewDB(map[any]map[any]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	transactions := db.Table["transactions"]

	firstTransaction := model.Transaction{
		ID:        uuid.New(),
		Name:      "Transaction 1",
		Email:     "transaction1@example.com",
		Password:  "hashedpassword1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	transactions[firstTransaction.ID.String()] = firstTransaction

	paginationReq := service.GetIssuesRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     3,
			PageSize: 5,
		},
	}

	result, totalCount, err := transactionService.GetIssues(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), totalCount)
}

func TestTransactionService_GetIssues_EmptyResult(t *testing.T) {
	db, err := database.NewDB(map[any]map[any]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	paginationReq := service.GetIssuesRequest{
		PaginationRequest: http_server.PaginationRequest{
			Page:     1,
			PageSize: 10,
		},
	}

	result, totalCount, err := transactionService.GetIssues(ctx, paginationReq)

	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), totalCount)
}
