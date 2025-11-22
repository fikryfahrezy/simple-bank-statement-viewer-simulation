package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *transactionService) GetBalance(ctx context.Context, id uuid.UUID) (BalanceResponse, error) {
	balance, err := s.transactionRepository.GetBalance(ctx, id)
	if err != nil {
		return BalanceResponse{}, err
	}

	response := ToBalanceResponse(balance)
	return response, nil
}
