package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
)

// Upload upload bank statement file
// @Summary Store a bank statement history
// @Description Storing data from bank statement file
// @Tags transaction
// @Accept json
// @Produce json
// @Param upload body service.UploadRequest true "Upload file request"
// @Success 201 {object} http_server.APIResponse{result=service.UploadResponse}
// @Failure 400 {object} http_server.APIResponse
// @Failure 500 {object} http_server.APIResponse
// @Router /upload [post]
func (h *TransactionHandler) Upload(w http.ResponseWriter, r *http.Request) {
	var req service.UploadRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.log.Error("Failed to bind request",
			slog.String("error", err.Error()),
		)
		http_server.BadRequestResponse(w, "Invalid request format", err)
		return
	}

	resp, err := h.transactionService.UploadStatement(r.Context(), req)
	if err != nil {
		h.translateServiceError(w, err, "Failed to upload statement")
		return
	}

	http_server.CreatedResponse(w, "Statement uploaded successfully", resp)
}
