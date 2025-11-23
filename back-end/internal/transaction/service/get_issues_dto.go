package service

import (
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
)

type IssueResponse struct {
	Timestamp   time.Time               `json:"timestamp"`
	Name        string                  `json:"name"`
	Type        model.TransactionType   `json:"type"`
	Amount      float64                 `json:"amount"`
	Status      model.TransactionStatus `json:"status"`
	Description string                  `json:"description"`
}

func ToIssuesResponse(u model.Transaction) IssueResponse {
	return IssueResponse{
		Timestamp:   time.Unix(u.Timestamp, 0).UTC(),
		Name:        u.Name,
		Type:        u.Type,
		Amount:      u.Amount,
		Status:      u.Status,
		Description: u.Description,
	}
}
