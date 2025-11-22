package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service/servicefakes"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionHandler_Upload_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	transactionID := uuid.New()
	expectedResponse := service.UploadResponse{
		ID:        transactionID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockService.UploadStatementReturns(expectedResponse, nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	requestBody := service.UploadRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	body, err := json.Marshal(requestBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Upload)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.UploadStatementCallCount())
	_, actualReq := mockService.UploadStatementArgsForCall(0)
	assert.Equal(t, requestBody.Name, actualReq.Name)
	assert.Equal(t, requestBody.Email, actualReq.Email)
	assert.Equal(t, requestBody.Password, actualReq.Password)
}
