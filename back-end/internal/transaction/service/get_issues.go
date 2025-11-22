package service

import (
	"context"
	"log/slog"
)

func (s *transactionService) GetIssues(ctx context.Context) ([]IssueResponse, error) {
	issues, err := s.transactionRepository.GetAllIssues(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]IssueResponse, len(issues))
	for i, issue := range issues {
		response := ToIssuesResponse(issue)
		responses[i] = response
	}

	s.log.Info("Issues listed successfully",
		slog.Int("count", len(responses)),
	)

	return responses, nil
}
