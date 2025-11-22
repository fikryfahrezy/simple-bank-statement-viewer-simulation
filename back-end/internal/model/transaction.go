package model

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "DEBIT"
	TransactionTypeCredit TransactionType = "CREDIT"
)

type TransactionStatus string

const (
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)

type Transaction struct {
	Timestamp   int64
	Name        string
	Type        TransactionType
	Amount      float64
	Status      TransactionStatus
	Description string
}
