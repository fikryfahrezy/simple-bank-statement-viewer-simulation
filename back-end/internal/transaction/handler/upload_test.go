package handler_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/logger"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/handler"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service/servicefakes"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_Upload_Success(t *testing.T) {
	mockService := &servicefakes.FakeTransactionService{}
	mockService.UploadStatementReturns(nil)

	transactionHandler := handler.NewTransactionHandler(logger.NewDiscardLogger(), mockService)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreatePart(
		map[string][]string{
			"Content-Disposition": {
				`form-data; name="file"; filename="statement.csv"`,
			},
			"Content-Type": {
				"text/csv",
			},
		},
	)
	assert.NoError(t, err)

	file := strings.NewReader("1624507883, JOHN DOE, DEBIT, 250000, SUCCESS, restaurant")
	_, err = io.Copy(part, file)
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionHandler.Upload)
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Verify service was called
	assert.Equal(t, 1, mockService.UploadStatementCallCount())
}
