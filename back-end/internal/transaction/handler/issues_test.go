package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service/servicefakes"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_Issues_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	expectedIssues := []service.IssueResponse{
		{
			Timestamp:   time.Date(2021, time.June, 24, 4, 11, 23, 0, time.UTC),
			Name:        "JOHN DOE",
			Type:        model.TransactionTypeDebit,
			Amount:      250000,
			Status:      model.TransactionStatusSuccess,
			Description: "restaurant",
		},
		{
			Timestamp:   time.Date(2021, time.June, 24, 4, 11, 23, 0, time.UTC),
			Name:        "COMPANY A",
			Type:        model.TransactionTypeCredit,
			Amount:      12000000,
			Status:      model.TransactionStatusFailed,
			Description: "salary",
		},
	}
	mockService.GetIssuesReturns(expectedIssues, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/issues", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Issues)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.GetIssuesCallCount())
}
