package usecases

import (
	"context"
	"fmt"
	"hashtracker/internal/entities"
)

type scanner struct {
	repo ETHTransactionRepository
}

func NewScanner(repo ETHTransactionRepository) ScannerUseCase {
	return &scanner{
		repo: repo,
	}
}

func (s *scanner) GetTransactions(ctx context.Context, address string, level int) (*entities.TransactionList, error) {
	fmt.Println("----> GetTransactions: ", address)
	transactions, getErr := s.repo.GetTransactions(ctx, address)
	if getErr != nil {
		return nil, getErr
	}
	return transactions, nil
}
