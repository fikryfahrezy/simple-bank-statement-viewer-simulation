package repository

import "github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/app_error"

// Repository errors
var (
	ErrTransactionsTableNotFound = app_error.New("FAILED_TO_GET_TRANSACTIONS", "failed to get transactions")
)
