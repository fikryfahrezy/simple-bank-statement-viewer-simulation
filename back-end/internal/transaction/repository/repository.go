package repository

import (
	"log/slog"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/database"
)

type transactionRepository struct {
	db  *database.DB
	log *slog.Logger
}

func NewTransactionRepository(log *slog.Logger, db *database.DB) *transactionRepository {
	return &transactionRepository{
		db:  db,
		log: log,
	}
}
