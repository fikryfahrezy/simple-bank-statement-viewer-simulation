package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service/servicefakes"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_Balance_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	transactionID := uuid.New()
	expectedResponse := service.BalanceResponse{
		ID:        transactionID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockService.GetBalanceReturns(expectedResponse, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/balance/{id}", nil)
	req.SetPathValue("id", transactionID.String())
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Balance)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with correct ID
	assert.Equal(t, 1, mockService.GetBalanceCallCount())
	_, actualID := mockService.GetBalanceArgsForCall(0)
	assert.Equal(t, transactionID, actualID)
}

func TestTransactionHandler_Balance_InvalidUUID(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/balance/{id}", nil)
	req.SetPathValue("id", "invalid-uuid")
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Balance)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Service should not be called on invalid UUID
	assert.Equal(t, 0, mockService.GetBalanceCallCount())
}
