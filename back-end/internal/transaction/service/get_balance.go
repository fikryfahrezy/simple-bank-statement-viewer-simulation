package service

import (
	"context"
)

func (s *transactionService) GetBalance(ctx context.Context) (BalanceResponse, error) {
	balance, err := s.transactionRepository.GetBalance(ctx)
	if err != nil {
		return BalanceResponse{}, err
	}

	response := BalanceResponse{
		Balance: balance,
	}
	return response, nil
}
