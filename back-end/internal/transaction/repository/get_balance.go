package repository

import (
	"context"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (r *transactionRepository) GetBalance(ctx context.Context) (float64, error) {
	trasactions, ok := r.db.Table["transactions"]
	if !ok {
		return 0, ErrTransactionsTableNotFound
	}

	var balance float64
	for _, row := range trasactions {
		transaction := row.(model.Transaction)
		if transaction.Status != model.TransactionStatusSuccess {
			continue
		}

		switch transaction.Type {
		case model.TransactionTypeCredit:
			balance += transaction.Amount
		case model.TransactionTypeDebit:
			balance -= transaction.Amount
		}
	}

	return balance, nil
}
