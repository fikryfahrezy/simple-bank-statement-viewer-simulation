package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (r *transactionRepository) Store(ctx context.Context, transaction model.Transaction) error {
	transactions, ok := r.db.Table["transactions"]
	if !ok {
		return ErrTransactionsTableNotFound
	}

	if err := r.db.BeginTx(); err != nil {
		r.log.Error("Failed to begin transaction", "error", err.Error())
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	transactions = append(transactions, transaction)
	r.db.Table["transactions"] = transactions

	fmt.Println(len(r.db.Table["transactions"]))

	if err := r.db.Commit(); err != nil {
		r.log.Error("Failed to commit transaction", "error", err.Error())
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	r.log.Info("Transaction created successfully",
		slog.Int64("transaction_timestamp", transaction.Timestamp),
	)

	return nil
}
