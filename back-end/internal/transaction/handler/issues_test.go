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

func TestTransactionHandler_Issues_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	expectedIssues := []service.IssueResponse{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Name:      "Jane Doe",
			Email:     "jane@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockService.GetIssuesReturns(expectedIssues, 2, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/issues", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Issues)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with default pagination
	assert.Equal(t, 1, mockService.GetIssuesCallCount())
	_, paginationReq := mockService.GetIssuesArgsForCall(0)
	assert.Equal(t, 1, paginationReq.Page)
	assert.Equal(t, 10, paginationReq.PageSize)
}

func TestTransactionHandler_Issues_WithPagination(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	expectedTransactions := []service.IssueResponse{
		{
			ID:        uuid.New(),
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	mockService.GetIssuesReturns(expectedTransactions, 1, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	req := httptest.NewRequest(http.MethodGet, "/issues?page=2&page_size=5", nil)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Issues)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	// Verify service was called with custom pagination
	assert.Equal(t, 1, mockService.GetIssuesCallCount())
	_, paginationReq := mockService.GetIssuesArgsForCall(0)
	assert.Equal(t, 2, paginationReq.Page)
	assert.Equal(t, 5, paginationReq.PageSize)
}
