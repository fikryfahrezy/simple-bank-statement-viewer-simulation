package handler

import (
	"log/slog"
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/google/uuid"
)

// Balance retrieves a total balance
// @Summary Get total balance
// @Description Retrieve total balance (credits – debits from successful transactions only)
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Balance ID"
// @Success 200 {object} http_server.APIResponse{result=service.BalanceResponse}
// @Failure 400 {object} http_server.APIResponse
// @Failure 404 {object} http_server.APIResponse
// @Failure 500 {object} http_server.APIResponse
// @Router /balance/{id} [get]
func (h *TransactionHandler) Balance(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.log.Warn("Invalid transaction ID parameter",
			slog.String("id", idParam),
		)
		http_server.BadRequestResponse(w, "Invalid balance UUID format", err)
		return
	}

	balanceResp, err := h.transactionService.GetBalance(r.Context(), id)
	if err != nil {
		h.translateServiceError(w, err, "Failed to get balance")
		return
	}

	http_server.SuccessResponse(w, "Balance retrieved successfully", balanceResp)
}
