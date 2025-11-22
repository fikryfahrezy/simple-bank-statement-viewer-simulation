package repository

import (
	"context"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	Store(ctx context.Context, transaction model.Transaction) error
	GetBalance(ctx context.Context, id uuid.UUID) (model.Transaction, error)
	GetIssues(ctx context.Context, limit, offset int) ([]model.Transaction, int64, error)
}
