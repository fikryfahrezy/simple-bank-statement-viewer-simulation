package repository

import (
	"context"
)

func (r *transactionRepository) GetBalance(ctx context.Context) (float64, error) {
	_, ok := r.db.Table["transactions"]
	if !ok {
		return 0, ErrTransactionsTableNotFound
	}

	return 0, nil
}
