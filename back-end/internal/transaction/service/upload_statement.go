package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

func (s *transactionService) UploadStatement(ctx context.Context, req UploadRequest) error {
	transactions, errorMap := ParseCSV(req.File)

	if len(errorMap) != 0 {
		s.log.Error("CSV validation fail",
			slog.Any("error", errorMap),
		)
		return &ParseError{Fields: errorMap}
	}

	if err := s.transactionRepository.Store(ctx, transactions); err != nil {
		return err
	}

	return nil
}

func ParseCSV(file io.Reader) ([]model.Transaction, map[string]any) {
	reader := csv.NewReader(file)

	keyPrefix := "line"
	errorMap := map[string]any{}
	transactions := []model.Transaction{}
	line := 0

	for {
		line++
		errorKey := createCSVErrorKey(keyPrefix, line)
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			errorMap[errorKey] = []string{
				"The content is invalid",
			}
			continue
		}

		if len(record) != 6 {
			errorMap[errorKey] = []string{
				fmt.Sprintf("Expected 6 columns, got %d", len(record[0])),
			}
			continue
		}

		errors := []string{}

		rawTimestamp := strings.TrimSpace(record[0])
		timestamp, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err != nil {
			errors = append(errors, "invalid 'timestamp', expected to be integer")
		}

		rawName := strings.TrimSpace(record[1])
		if rawName == "" {
			errors = append(errors, "'name' is required")
		}

		rawType := model.TransactionType(strings.TrimSpace(record[2]))
		if !rawType.Valid() {
			errors = append(errors, fmt.Sprintf("The posible value for 'type' is %s, %s", model.TransactionTypeDebit, model.TransactionTypeCredit))
		}

		rawAmount := strings.TrimSpace(record[3])
		amount, err := strconv.ParseFloat(rawAmount, 64)
		if err != nil {
			errors = append(errors, "invalid 'amount', expected to be number")
		}

		rawStatus := model.TransactionStatus(strings.TrimSpace(record[4]))
		if !rawStatus.Valid() {
			errors = append(errors, fmt.Sprintf("The posible value for 'status' is %s, %s, %s", model.TransactionStatusSuccess, model.TransactionStatusPending, model.TransactionStatusFailed))
		}

		rawDescription := strings.TrimSpace(record[5])
		if rawName == "" {
			errors = append(errors, "'description' is required")
		}

		if len(errors) == 0 {
			transactions = append(transactions, model.Transaction{
				Timestamp:   timestamp,
				Name:        rawName,
				Type:        rawType,
				Amount:      amount,
				Status:      rawStatus,
				Description: rawDescription,
			})
		} else {
			errorMap[errorKey] = errors
		}
	}

	return transactions, errorMap
}

func createCSVErrorKey(prefix string, line int) string {
	return fmt.Sprintf("%s[%d]", prefix, line)
}
