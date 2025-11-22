package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/google/uuid"
)

func (r *transactionRepository) Store(ctx context.Context, transaction model.Transaction) error {
	transactions, ok := r.db.Table["transactions"]
	if !ok {
		return ErrTransactionsTableNotFound
	}

	now := time.Now()
	transaction.CreatedAt = now
	transaction.UpdatedAt = now

	// Generate UUIDv7 for the transaction ID
	transactionID := uuid.Must(uuid.NewV7())
	transaction.ID = transactionID

	if err := r.db.BeginTx(); err != nil {
		r.log.Error("Failed to begin transaction", "error", err.Error())
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	transactions[transactionID.String()] = transaction

	if err := r.db.Commit(); err != nil {
		r.log.Error("Failed to commit transaction", "error", err.Error())
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	r.log.Info("Transaction created successfully",
		slog.String("transaction_id", transactionID.String()),
		slog.String("email", transaction.Email),
	)

	return nil
}
