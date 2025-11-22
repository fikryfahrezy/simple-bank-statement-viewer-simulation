package repository

import (
	"context"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (r *transactionRepository) GetIssues(ctx context.Context, limit, offset int) ([]model.Transaction, int64, error) {
	transactionTable, ok := r.db.Table["transactions"]
	if !ok {
		return []model.Transaction{}, 0, ErrTransactionsTableNotFound
	}

	var transactions []model.Transaction
	for _, row := range transactionTable {
		transaction, ok := row.(model.Transaction)
		if !ok {
			continue
		}
		transactions = append(transactions, transaction)
	}

	return transactions, int64(len(transactions)), nil
}
