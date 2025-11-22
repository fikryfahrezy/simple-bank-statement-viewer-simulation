package service

import (
	"context"
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (s *transactionService) UploadStatement(ctx context.Context, req UploadRequest) error {
	transaction := model.Transaction{
		Timestamp:   1624507883,
		Name:        "JOHN DOE",
		Type:        model.TransactionTypeDebit,
		Amount:      250000,
		Status:      model.TransactionStatusSuccess,
		Description: "restaurant",
	}

	if err := s.transactionRepository.Store(ctx, transaction); err != nil {
		return err
	}

	s.log.Info("Transaction created successfully",
		slog.Int64("transaction_timestamp", transaction.Timestamp),
	)

	return nil
}
