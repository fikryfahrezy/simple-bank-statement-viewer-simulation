package handler

import (
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
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
// @Router /balance [get]
func (h *TransactionHandler) Balance(w http.ResponseWriter, r *http.Request) {
	balanceResp, err := h.transactionService.GetBalance(r.Context())
	if err != nil {
		h.translateServiceError(w, err, "Failed to get balance")
		return
	}

	http_server.SuccessResponse(w, "Balance retrieved successfully", balanceResp)
}
