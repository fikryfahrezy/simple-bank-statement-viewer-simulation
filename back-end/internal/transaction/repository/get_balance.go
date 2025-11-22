package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/google/uuid"
)

func (r *transactionRepository) GetBalance(ctx context.Context, id uuid.UUID) (model.Transaction, error) {
	transactions, ok := r.db.Table["transactions"]
	if !ok {
		return model.Transaction{}, ErrTransactionsTableNotFound
	}

	transactionItem, ok := transactions[id.String()]
	if !ok {
		return model.Transaction{}, fmt.Errorf("transaction not found")
	}

	transaction, ok := transactionItem.(model.Transaction)
	if !ok {
		r.log.Error("Failed to get transaction by ID",
			slog.String("transaction_id", id.String()),
		)
		return model.Transaction{}, fmt.Errorf("failed to case transaction object")
	}

	return transaction, nil
}
