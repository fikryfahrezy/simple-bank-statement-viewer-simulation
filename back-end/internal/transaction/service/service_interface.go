package service

//counterfeiter:generate -o servicefakes/fake_transaction_service.go . TransactionService

import (
	"context"

	"github.com/google/uuid"
)

type TransactionService interface {
	UploadStatement(ctx context.Context, req UploadRequest) (UploadResponse, error)
	GetBalance(ctx context.Context, id uuid.UUID) (BalanceResponse, error)
	GetIssues(ctx context.Context, req GetIssuesRequest) ([]IssueResponse, int64, error)
}
