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
		Status:      model.TransactionStatusSuccess,
		Description: "salary",
	}
	thirdTransaction := model.Transaction{
		Timestamp:   1624608050,
		Name:        "E-COMMERCE A",
		Type:        model.TransactionTypeDebit,
		Amount:      150000,
		Status:      model.TransactionStatusFailed,
		Description: "clothes",
	}

	transactions := db.Table["transactions"]
	transactions = append(transactions, firstTransaction)
	transactions = append(transactions, secondTransaction)
	transactions = append(transactions, thirdTransaction)
	db.Table["transactions"] = transactions

	result, err := transactionService.GetBalance(ctx)

	assert.NoError(t, err)
	assert.Equal(t, float64(11750000), result.Balance)
}
