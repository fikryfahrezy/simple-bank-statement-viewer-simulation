package repository

import (
	"context"
	"fmt"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (r *transactionRepository) Store(ctx context.Context, transactions []model.Transaction) error {
	transactionsTable, ok := r.db.Table["transactions"]
	if !ok {
		return ErrTransactionsTableNotFound
	}

	if err := r.db.BeginTx(); err != nil {
		r.log.Error("Failed to begin transaction", "error", err.Error())
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	for _, transaction := range transactions {
		transactionsTable = append(transactionsTable, transaction)
	}
	r.db.Table["transactions"] = transactionsTable

	if err := r.db.Commit(); err != nil {
		r.log.Error("Failed to commit transaction", "error", err.Error())
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
