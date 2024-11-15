package usecases

import (
	"context"
	"hashtracker/internal/entities"
)

type ETHTransactionRepository interface {
	// GetTransactions retrieves all transactions on Ethereum main net
	// associated to the specified address
	GetTransactions(ctx context.Context, address string) (*entities.TransactionList, error)

	// GetOutTransactions get just the outgoing transactions associated to a wallet.
	GetOutTransactions(ctx context.Context, address string) (*entities.TransactionList, error)
}

type ScannerUseCase interface {
	// GetTransactions all transactions on Ethereum main net associated to the specified address
	GetTransactions(ctx context.Context, address string, level int) (*entities.TransactionList, error)
}
