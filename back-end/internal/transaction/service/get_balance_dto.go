package service

import (
	"time"

	"github.com/fikryfahrezy/simple-bank-statement-viewer-simulation/internal/model"
	"github.com/google/uuid"
)

type BalanceResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToBalanceResponse(u model.Transaction) BalanceResponse {
	return BalanceResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
