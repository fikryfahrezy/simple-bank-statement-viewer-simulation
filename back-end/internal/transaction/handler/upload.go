package handler

import (
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
)

// Upload upload bank statement file
// @Summary Store a bank statement history
// @Description Storing data from bank statement file
// @Tags transactions
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file (only file supported)"
// @Success 201 {object} http_server.APIResponse{result=nil}
// @Failure 400 {object} http_server.APIResponse
// @Failure 500 {object} http_server.APIResponse
// @Router /upload [post]
func (h *TransactionHandler) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // 32 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			h.log.Error("Failed to close the file",
				slog.String("error", err.Error()))
		}
	}()

	if header.Filename == "" {
		http_server.BadRequestResponse(w, "no file provided", err)
		return
	}

	if header.Header.Get("Content-Type") != "text/csv" {
		http_server.BadRequestResponse(w, "file expected to be csv", err)
		return
	}

	req := service.UploadRequest{
		File: file,
	}

	err = h.transactionService.UploadStatement(r.Context(), req)
	if err != nil {
		if ve, ok := err.(*service.ParseError); ok {
			http_server.ValidationErrorResponse(w, "Validation failed", ve.Fields)
			return
		}
		h.translateServiceError(w, err, "Failed to upload statement")
		return
	}

	http_server.CreatedResponse(w, "Statement uploaded successfully", nil)
}
