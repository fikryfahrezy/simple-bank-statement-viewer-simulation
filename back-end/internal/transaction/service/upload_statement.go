package service

import (
	"context"
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *transactionService) UploadStatement(ctx context.Context, req UploadRequest) (UploadResponse, error) {
	s.log.Info("Uploading bank statement",
		slog.String("email", req.Email),
		slog.String("name", req.Name),
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed to hash password",
			slog.String("error", err.Error()),
		)
		return UploadResponse{}, err
	}

	transaction := model.Transaction{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.transactionRepository.Store(ctx, transaction); err != nil {
		return UploadResponse{}, err
	}

	response := ToUploadResponse(transaction)
	s.log.Info("Transaction created successfully",
		slog.String("transaction_id", transaction.ID.String()),
		slog.String("email", transaction.Email),
	)

	return response, nil
}
