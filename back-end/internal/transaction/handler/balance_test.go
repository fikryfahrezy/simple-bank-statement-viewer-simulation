package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service/servicefakes"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_Balance_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	expectedResponse := service.BalanceResponse{
		Balance: 0,
	}
	mockService.GetBalanceReturns(expectedResponse, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/balance", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Balance)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.GetBalanceCallCount())
}
