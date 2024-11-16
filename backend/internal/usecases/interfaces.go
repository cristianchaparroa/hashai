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

type ENSRepository interface {
	Resolve(ctx context.Context, name string) (*entities.ENSResponse, error)
}

type PolygonRepository interface {
	Resolve(ctx context.Context, address string) (*entities.PolygonResponse, error)
}

type ScannerUseCase interface {
	// GetTransactions all transactions on Ethereum main net associated to the specified address
	GetTransactions(ctx context.Context, address string, level int) (*entities.TransactionList, error)

	Scan(ctx context.Context, address string, level int) (*entities.TransactionList, error)
}

type ENSValidator interface {
	IsValid(ctx context.Context, address string) bool
}
