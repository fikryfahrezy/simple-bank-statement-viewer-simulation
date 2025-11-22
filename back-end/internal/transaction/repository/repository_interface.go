package repository

import (
	"context"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

type TransactionRepository interface {
	Store(ctx context.Context, transaction model.Transaction) error
	GetBalance(ctx context.Context) (float64, error)
	GetIssues(ctx context.Context, limit, offset int) ([]model.Transaction, int64, error)
}
