package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_GetBalance_Success(t *testing.T) {
	db, err := database.NewDB(map[any]map[any]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionsRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionsRepo)
	ctx := context.Background()

	transactionID := uuid.New()
	expectedTransaction := model.Transaction{
		ID:        transactionID,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	transactions := db.Table["transactions"]
	transactions[transactionID.String()] = expectedTransaction

	result, err := transactionService.GetBalance(ctx, transactionID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction.ID, result.ID)
	assert.Equal(t, expectedTransaction.Name, result.Name)
	assert.Equal(t, expectedTransaction.Email, result.Email)
	assert.Equal(t, expectedTransaction.CreatedAt, result.CreatedAt)
	assert.Equal(t, expectedTransaction.UpdatedAt, result.UpdatedAt)
}
