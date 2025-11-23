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
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_GetIssues_Success(t *testing.T) {
	firstTransaction := model.Transaction{
		Timestamp:   1624507883,
		Name:        "JOHN DOE",
		Type:        model.TransactionTypeDebit,
		Amount:      250000,
		Status:      model.TransactionStatusSuccess,
		Description: "restaurant",
	}
	secondTransaction := model.Transaction{
		Timestamp:   1624608050,
		Name:        "COMPANY A",
		Type:        model.TransactionTypeCredit,
		Amount:      12000000,
		Status:      model.TransactionStatusFailed,
		Description: "salary",
	}

	db, err := database.NewDB(map[any][]any{
		"transactions": {firstTransaction, secondTransaction},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	result, err := transactionService.GetIssues(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)

	// Verify first transaction
	assert.Equal(t, time.Time(time.Date(2021, time.June, 24, 4, 11, 23, 0, time.UTC)), result[0].Timestamp)
	assert.Equal(t, firstTransaction.Name, result[0].Name)
	assert.Equal(t, firstTransaction.Type, result[0].Type)
	assert.Equal(t, firstTransaction.Amount, result[0].Amount)
	assert.Equal(t, firstTransaction.Status, result[0].Status)
	assert.Equal(t, firstTransaction.Description, result[0].Description)

	// Verify second transaction
	assert.Equal(t, time.Time(time.Date(2021, time.June, 25, 8, 0, 50, 0, time.UTC)), result[1].Timestamp)
	assert.Equal(t, secondTransaction.Name, result[1].Name)
	assert.Equal(t, secondTransaction.Type, result[1].Type)
	assert.Equal(t, secondTransaction.Amount, result[1].Amount)
	assert.Equal(t, secondTransaction.Status, result[1].Status)
	assert.Equal(t, secondTransaction.Description, result[1].Description)
}

func TestTransactionService_GetIssues_EmptyResult(t *testing.T) {
	db, err := database.NewDB(map[any][]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	result, err := transactionService.GetIssues(ctx)

	assert.NoError(t, err)
	assert.Empty(t, result)
}
