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

	text := "1624507883, JOHN DOE, DEBIT, 250000, SUCCESS, restaurant"
	text += "\n1624608050, E-COMMERCE A, DEBIT, 150000, FAILED, clothes"
	text += "\n1624512883, COMPANY A, CREDIT, 12000000, SUCCESS, salary"
	text += "\n1624615065, E-COMMERCE B, DEBIT, 150000, PENDING, clothes"

	req := service.UploadRequest{
		File: strings.NewReader(text),
	}

	prevTransactions := db.Table["transactions"]
	assert.Equal(t, 0, len(prevTransactions))

	err = transactionService.UploadStatement(ctx, req)
	assert.NoError(t, err)

	newTransactions := db.Table["transactions"]
	assert.Equal(t, 4, len(newTransactions))

	var actualFirstTransaction model.Transaction
	actualFirstTransaction = newTransactions[0].(model.Transaction)

	assert.Equal(t, int64(1624507883), actualFirstTransaction.Timestamp)
	assert.Equal(t, "JOHN DOE", actualFirstTransaction.Name)
	assert.Equal(t, model.TransactionTypeDebit, actualFirstTransaction.Type)
	assert.Equal(t, float64(250000), actualFirstTransaction.Amount)
	assert.Equal(t, model.TransactionStatusSuccess, actualFirstTransaction.Status)
	assert.Equal(t, "restaurant", actualFirstTransaction.Description)

	var actualSecondTransaction model.Transaction
	actualSecondTransaction = newTransactions[1].(model.Transaction)

	assert.Equal(t, int64(1624608050), actualSecondTransaction.Timestamp)
	assert.Equal(t, "E-COMMERCE A", actualSecondTransaction.Name)
	assert.Equal(t, model.TransactionTypeDebit, actualSecondTransaction.Type)
	assert.Equal(t, float64(150000), actualSecondTransaction.Amount)
	assert.Equal(t, model.TransactionStatusFailed, actualSecondTransaction.Status)
	assert.Equal(t, "clothes", actualSecondTransaction.Description)

	var actualThirdTransaction model.Transaction
	actualThirdTransaction = newTransactions[2].(model.Transaction)

	assert.Equal(t, int64(1624512883), actualThirdTransaction.Timestamp)
	assert.Equal(t, "COMPANY A", actualThirdTransaction.Name)
	assert.Equal(t, model.TransactionTypeCredit, actualThirdTransaction.Type)
	assert.Equal(t, float64(12000000), actualThirdTransaction.Amount)
	assert.Equal(t, model.TransactionStatusSuccess, actualThirdTransaction.Status)
	assert.Equal(t, "salary", actualThirdTransaction.Description)

	var actualFourthTransaction model.Transaction
	actualFourthTransaction = newTransactions[3].(model.Transaction)

	assert.Equal(t, int64(1624615065), actualFourthTransaction.Timestamp)
	assert.Equal(t, "E-COMMERCE B", actualFourthTransaction.Name)
	assert.Equal(t, model.TransactionTypeDebit, actualFourthTransaction.Type)
	assert.Equal(t, float64(150000), actualFourthTransaction.Amount)
	assert.Equal(t, model.TransactionStatusPending, actualFourthTransaction.Status)
	assert.Equal(t, "clothes", actualFourthTransaction.Description)
}

func TestTransactionService_CreateTransaction_ParseCSV_AllLines_Failed(t *testing.T) {
	// Setup
	db, err := database.NewDB(map[any][]any{
		"transactions": {},
	})

	log := logger.NewDiscardLogger()
	assert.NoError(t, err)
	transactionRepo := repository.NewTransactionRepository(log, db)
	transactionService := service.NewTransactionService(log, transactionRepo)
	ctx := context.Background()

	text := "1624507883, JOHN DOE, UNKNOWN, 250000, SUCCESS, restaurant"
	text += "\nWIB, E-COMMERCE A, DEBIT, 150000, FAILED, clothes"
	text += "\n1624512883, COMPANY A, CREDIT, DUA-BELAS, SUCCESS, salary"
	text += "\n1624615065, E-COMMERCE B, DEBIT, 150000, UNKNOWN, clothes\n"
	text += "\n1624615066,, DEBIT, 150000, SUCCESS, clothes\n"
	text += "\n1624615066, E-COMMERCE C, DEBIT, 150000, SUCCESS,\n"

	req := service.UploadRequest{
		File: strings.NewReader(text),
	}

	prevTransactions := db.Table["transactions"]
	assert.Equal(t, 0, len(prevTransactions))

	err = transactionService.UploadStatement(ctx, req)
	ve, ok := err.(*service.ParseError)
	assert.Equal(t, true, ok)

	assert.Equal(t, 6, len(ve.Fields))

	lineOneErrorMessages, ok := ve.Fields["line[1]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "The posible value for 'type' is DEBIT, CREDIT", lineOneErrorMessages[0])

	lineTwoErrorMessages, ok := ve.Fields["line[2]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "Invalid 'timestamp', expected to be integer", lineTwoErrorMessages[0])

	lineThreeErrorMessages, ok := ve.Fields["line[3]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "Invalid 'amount', expected to be number", lineThreeErrorMessages[0])

	lineFourErrorMessages, ok := ve.Fields["line[4]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "The posible value for 'status' is SUCCESS, PENDING, FAILED", lineFourErrorMessages[0])

	lineFiveErrorMessages, ok := ve.Fields["line[5]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "The 'name' is required", lineFiveErrorMessages[0])

	lineSixErrorMessages, ok := ve.Fields["line[6]"].([]string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "The 'description' is required", lineSixErrorMessages[0])

	newTransactions := db.Table["transactions"]
	assert.Equal(t, 0, len(newTransactions))
}
