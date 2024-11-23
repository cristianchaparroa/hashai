package usecases

import (
	"context"
	"hashtracker/internal/entities/blockscout"
	"hashtracker/internal/entities/polygon"
	"hashtracker/internal/entities/thegraph"
)

type ETHTransactionRepository interface {
	// GetTransactions retrieves all transactions on Ethereum main net
	// associated to the specified address
	GetTransactions(ctx context.Context, address string) (*blockscout.TransactionList, error)

	// GetOutTransactions get just the outgoing transactions associated to a wallet.
	GetOutTransactions(ctx context.Context, address string) (*blockscout.TransactionList, error)
}

type ENSRepository interface {
	Resolve(ctx context.Context, name string) (*thegraph.ENSResponse, error)
}

type PolygonRepository interface {
	CreateReport(ctx context.Context, req *polygon.ReportRequest) (*polygon.PolygonResponse, error)
}

type ScannerUseCase interface {
	// GetTransactions all transactions on Ethereum main net associated to the specified address
	GetTransactions(ctx context.Context, address string, level int) (*blockscout.TransactionList, error)

	Scan(ctx context.Context, address string, level int) (*blockscout.TransactionList, error)
}

type ENSValidator interface {
	IsValid(ctx context.Context, address string) bool
}
