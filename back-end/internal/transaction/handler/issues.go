package handler

import (
	"net/http"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
)

// Issues retrieves a list of non-successful transactions
// @Summary List non-successful transactions
// @Description Retrieve a paginated list of non-successful transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Success 200 {object} http_server.ListAPIResponse{result=[]service.IssueResponse}
// @Failure 500 {object} http_server.APIResponse
// @Router /issues [get]
func (h *TransactionHandler) Issues(w http.ResponseWriter, r *http.Request) {
	issues, err := h.transactionService.GetIssues(r.Context())
	if err != nil {
		h.translateServiceError(w, err, "Failed to list issues")
		return
	}

	http_server.SuccessResponse(w, "Issues retrieved successfully", issues)
}
