package polygon

import (
	"context"
	"fmt"
	"hashtracker/config"
	"hashtracker/internal/entities/polygon"
	"hashtracker/internal/usecases"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type hashReportRepository struct {
	contract *ContractConfig
}

func NewHashReportRepository(cfg *config.Config) usecases.HashReportRepository {
	contract := &ContractConfig{
		PrivateKey:      cfg.Polygon.PrivateKey,
		RpcURL:          cfg.Polygon.RpcURl,
		ChainID:         cfg.Polygon.ChainID,
		AbiFilePath:     cfg.ReportContract.AbiFile,
		ContractAddress: cfg.ReportContract.Address,
	}
	return &hashReportRepository{
		contract: contract,
	}
}
func (p *hashReportRepository) CreateReport(ctx context.Context, req *polygon.ReportRequest) (*polygon.PolygonResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, p.contract.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Polygon node: %w", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(p.contract.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(p.contract.ChainID))
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Contract address
	contractAddress := common.HexToAddress(p.contract.ContractAddress)

	// Load ABI from a file
	abiContent, err := os.ReadFile(p.contract.AbiFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %w", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiContent)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Pack arguments for createReport
	reportedAddress := common.HexToAddress(req.Address)
	category := big.NewInt(1) // Example: Phishing

	txData, err := parsedABI.Pack("createReport",
		reportedAddress,
		category,
		req.Comments,
		req.Source,
		req.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack arguments for createReport: %w", err)
	}

	// Get current nonce
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch nonce: %w", err)
	}

	// Add more gas headroom for the estimation
	gasLimit := uint64(300000) // Set a reasonable fixed gas limit

	// Get gas price with some headroom
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch gas price: %w", err)
	}

	// Add 20% to the gas price for better chances of transaction success
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

	// Construct transaction
	tx := types.NewTransaction(
		nonce,
		contractAddress,
		big.NewInt(0),
		gasLimit,
		gasPrice,
		txData,
	)

	// Sign transaction
	chainID := big.NewInt(p.contract.ChainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return &polygon.PolygonResponse{
		HashTransaction: signedTx.Hash().Hex(),
	}, nil
}
