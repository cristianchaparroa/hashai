package polygon

import (
	"context"
	"fmt"
	"hashtracker/internal/contracts"
	"hashtracker/internal/entities/polygon"
	"hashtracker/internal/usecases"
	"hashtracker/pkg/eth"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

type BlacklistNoopRepository struct {
}

func NewBlacklistNoopRepository() (usecases.BlacklistReportRepository, error) {
	return &BlacklistNoopRepository{}, nil
}

func (r *BlacklistNoopRepository) CreateBatchReport(ctx context.Context, reports []*polygon.BatchReport) error {
	return nil
}

type BlacklistRepository struct {
	client   eth.Client
	config   *ContractConfig
	contract *contracts.Blacklist
}

func NewBlacklistRepository(config *ContractConfig, client eth.Client) (usecases.BlacklistReportRepository, error) {
	if client == nil {
		return nil, fmt.Errorf("ethclient is nil")
	}

	address := config.GetContractAddress()
	contract, err := contracts.NewBlacklist(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Verify contract by calling a view function
	_, err = contract.Nonces(&bind.CallOpts{}, address)
	if err != nil {
		return nil, fmt.Errorf("contract verification failed: %w", err)
	}

	return &BlacklistRepository{
		contract: contract,
		client:   client,
		config:   config,
	}, nil
}

func (r *BlacklistRepository) CreateBatchReport(ctx context.Context, reports []*polygon.BatchReport) error {
	const batchSize = 100 // Adjust based on gas limits and network conditions

	chunks := len(reports)/batchSize + 1
	for i := 0; i < chunks; i++ {
		start := i * batchSize
		end := start + batchSize
		if end > len(reports) {
			end = len(reports)
		}

		if start >= len(reports) {
			break
		}

		err := r.processBatchReport(ctx, reports[start:end])
		if err != nil {
			return fmt.Errorf("chunk %d failed: %w", i, err)
		}

		// Add delay between chunks to avoid nonce issues
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func (r *BlacklistRepository) processBatchReport(ctx context.Context, reports []*polygon.BatchReport) error {
	if len(reports) == 0 {
		return nil
	}

	auth, err := r.getTransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transactor error: %w", err)
	}

	auth.GasLimit = 500000 * uint64(len(reports)) // Scale gas limit with batch size
	auth.GasTipCap = big.NewInt(30000000000)
	auth.GasFeeCap = big.NewInt(50000000000)

	privateKey, _ := r.config.GetECDSAPrivateKey()
	userOps := make([]contracts.BlacklistUserOperation, len(reports))

	nonce, err := r.contract.Nonces(&bind.CallOpts{}, auth.From)
	if err != nil {
		return fmt.Errorf("nonce error: %w", err)
	}

	for i, report := range reports {
		userOps[i] = contracts.BlacklistUserOperation{
			Sender:          auth.From,
			ReportedAddress: report.ReportedAddress,
			Category:        big.NewInt(1),
			Comments:        report.Comments,
			Source:          report.Source,
			Date:            big.NewInt(time.Now().Unix()),
			Nonce:           new(big.Int).Add(nonce, big.NewInt(int64(i))),
			Signature:       []byte{},
		}

		hash, err := r.contract.GetOperationHash(&bind.CallOpts{}, userOps[i])
		if err != nil {
			return fmt.Errorf("hash error for operation %d: %w", i, err)
		}

		signature, err := crypto.Sign(hash[:], privateKey)
		if err != nil {
			return fmt.Errorf("signing error for operation %d: %w", i, err)
		}
		signature[64] += 27
		userOps[i].Signature = signature
	}

	tx, err := r.contract.ExecuteBatchOperations(auth, userOps)
	if err != nil {
		return fmt.Errorf("batch execution error: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, r.client, tx)
	if err != nil {
		return fmt.Errorf("mining error: %w", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	return nil
}

func (r *BlacklistRepository) getTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	chainID, err := r.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	pk, pkErr := r.config.GetECDSAPrivateKey()
	if pkErr != nil {
		return nil, pkErr
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
