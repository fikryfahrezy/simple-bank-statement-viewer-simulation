package service

import (
	"context"
	"log/slog"
)

func (s *transactionService) GetIssues(ctx context.Context, req GetIssuesRequest) ([]IssueResponse, int64, error) {
	s.log.Info("Listing issues",
		slog.Int("page", req.Page),
		slog.Int("page_size", req.PageSize),
	)

	offset := (req.Page - 1) * req.PageSize

	issues, total, err := s.transactionRepository.GetIssues(ctx, req.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}

	var responses []IssueResponse
	for _, issue := range issues {
		response := ToIssuesListResponse(issue)
		responses = append(responses, response)
	}

	s.log.Info("Issues listed successfully",
		slog.Int("count", len(responses)),
		slog.Int64("total", total),
	)

	return responses, total, nil
}
