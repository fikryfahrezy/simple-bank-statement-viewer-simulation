package service_test

import (
	"context"
	"testing"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_GetBalance_Success(t *testing.T) {
	db, err := database.NewDB(map[any][]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionsRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionsRepo)
	ctx := context.Background()

	expectedTransaction := model.Transaction{
		Timestamp:   1624507883,
		Name:        "JOHN DOE",
		Type:        model.TransactionTypeDebit,
		Amount:      250000,
		Status:      model.TransactionStatusSuccess,
		Description: "restaurant",
	}

	transactions := db.Table["transactions"]
	transactions = append(transactions, expectedTransaction)
	db.Table["transactions"] = transactions

	result, err := transactionService.GetBalance(ctx)

	assert.NoError(t, err)
	assert.Equal(t, float64(0), result.Balance)
}
