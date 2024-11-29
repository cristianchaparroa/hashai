package usecases

import (
	"context"
	"hashtracker/internal/entities/blockscout"
)

type ScannerUseCase interface {
	// GetTransactions all transactions on Ethereum main net associated to the specified address
	GetTransactions(ctx context.Context, address string, level int) (*blockscout.TransactionList, error)

	Scan(ctx context.Context, address string, level int) (*blockscout.TransactionList, error)
}

type ENSValidator interface {
	IsValid(ctx context.Context, address string) bool
}
