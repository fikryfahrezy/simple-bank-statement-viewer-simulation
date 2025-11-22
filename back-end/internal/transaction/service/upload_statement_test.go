package service_test

import (
	"context"
	"strings"
	"testing"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_CreateTransaction_Success(t *testing.T) {
	// Setup
	db, err := database.NewDB(map[any][]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	req := service.UploadRequest{
		File: strings.NewReader("1624507883, JOHN DOE, DEBIT, 250000, SUCCESS, restaurant"),
	}

	prevTransactions := db.Table["transactions"]
	assert.Equal(t, 0, len(prevTransactions))

	err = transactionService.UploadStatement(ctx, req)
	assert.NoError(t, err)

	newTransactions := db.Table["transactions"]
	assert.Equal(t, 1, len(newTransactions))

	var actualTransaction model.Transaction
	for _, transaction := range newTransactions {
		actualTransaction = transaction.(model.Transaction)
	}

	assert.Equal(t, int64(1624507883), actualTransaction.Timestamp)
	assert.Equal(t, "JOHN DOE", actualTransaction.Name)
	assert.Equal(t, model.TransactionTypeDebit, actualTransaction.Type)
	assert.Equal(t, float64(250000), actualTransaction.Amount)
	assert.Equal(t, model.TransactionStatusSuccess, actualTransaction.Status)
	assert.Equal(t, "restaurant", actualTransaction.Description)
}
