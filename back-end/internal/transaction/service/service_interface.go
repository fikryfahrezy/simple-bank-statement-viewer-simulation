package service

//counterfeiter:generate -o servicefakes/fake_transaction_service.go . TransactionService

import (
	"context"
)

type TransactionService interface {
	UploadStatement(ctx context.Context, req UploadRequest) error
	GetBalance(ctx context.Context) (BalanceResponse, error)
	GetIssues(ctx context.Context, req GetIssuesRequest) ([]IssueResponse, int64, error)
}
