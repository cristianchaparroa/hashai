package usecases

import (
	"context"
	"fmt"
	"hashtracker/internal/entities"
)

type scanner struct {
	repo         ETHTransactionRepository
	ensRepo      ENSRepository
	ensValidator ENSValidator
}

func NewScanner(
	repo ETHTransactionRepository,
	ensRepo ENSRepository,
	ensValidator ENSValidator,
) ScannerUseCase {
	return &scanner{
		repo:         repo,
		ensValidator: ensValidator,
		ensRepo:      ensRepo,
	}
}

func (s *scanner) Scan(ctx context.Context, address string, level int) (*entities.TransactionList, error) {
	isValidENS := s.ensValidator.IsValid(ctx, address)
	fmt.Println("--> Is valid ENS: ", isValidENS)
	if isValidENS {
		result, err := s.ensRepo.Resolve(ctx, address)
		if err != nil {
			return nil, err
		}
		address = result.Address
	}

	return s.GetTransactions(ctx, address, level)
}

func (s *scanner) GetTransactions(ctx context.Context, address string, level int) (*entities.TransactionList, error) {
	fmt.Println("----> GetTransactions: ", address)
	transactions, getErr := s.repo.GetTransactions(ctx, address)
	if getErr != nil {
		return nil, getErr
	}
	return transactions, nil
}
