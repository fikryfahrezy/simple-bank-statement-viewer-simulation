package service

import (
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/transaction/repository"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
	log                   *slog.Logger
}

func NewTransactionService(log *slog.Logger, transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		log:                   log,
	}
}
