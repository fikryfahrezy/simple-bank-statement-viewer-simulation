package repository

import (
	"context"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (r *transactionRepository) GetAllIssues(ctx context.Context) ([]model.Transaction, error) {
	transactionTable, ok := r.db.Table["transactions"]
	if !ok {
		return []model.Transaction{}, ErrTransactionsTableNotFound
	}

	var transactions []model.Transaction
	for _, row := range transactionTable {
		transaction, ok := row.(model.Transaction)
		if !ok {
			continue
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
