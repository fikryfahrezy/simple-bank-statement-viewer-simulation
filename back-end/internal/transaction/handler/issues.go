package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/http_server"
	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/service"
)

// Issues retrieves a list of non-successful transactions
// @Summary List non-successful transactions
// @Description Retrieve a paginated list of non-successful transactions
// @Tags transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Success 200 {object} http_server.ListAPIResponse{result=[]service.IssueResponse}
// @Failure 500 {object} http_server.APIResponse
// @Router /issues [get]
func (h *TransactionHandler) Issues(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pageParam := queryParams.Get("page")
	pageSizeParam := queryParams.Get("page_size")

	page := 1
	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	pageSize := 10
	if pageSizeParam != "" {
		if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	paginationReq := http_server.PaginationRequest{
		Page:     page,
		PageSize: pageSize,
	}
	issues, totalCount, err := h.transactionService.GetIssues(r.Context(), service.GetIssuesRequest{
		PaginationRequest: paginationReq,
	})
	if err != nil {
		h.translateServiceError(w, err, "Failed to list issues")
		return
	}

	totalPages := int64(math.Ceil(float64(totalCount) / float64(pageSize)))
	pagination := http_server.CreatePaginationResponse(totalCount, totalPages, page, pageSize)

	http_server.ListSuccessResponse(w, "Issues retrieved successfully", issues, pagination)
}
