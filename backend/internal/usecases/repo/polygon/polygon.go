package polygon

import (
	"context"
	"fmt"
	"hashtracker/config"
	"hashtracker/internal/entities"
	"hashtracker/internal/usecases"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type polygonRepository struct {
	privateKey      string
	rpcURL          string
	chainID         int64
	abiFilePath     string
	contractAddress string
}

func NewPolygonRepository(abiFilePath string, contractAddress string, chainID int64, cfg config.Config) usecases.PolygonRepository {
	return &polygonRepository{
		privateKey:      cfg.PolygonPrivateKey,
		rpcURL:          cfg.PolygonRpcURl,
		chainID:         chainID,
		abiFilePath:     abiFilePath,
		contractAddress: contractAddress,
	}
}
func (p *polygonRepository) Resolve(ctx context.Context, address string) (*entities.PolygonResponse, error) {
	// Connect to Polygon Mumbai
	client, err := ethclient.Dial(p.rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Polygon node: %w", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(p.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(p.chainID))
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Contract address
	contractAddress := common.HexToAddress(p.contractAddress)

	// Load ABI from a file
	abiContent, err := os.ReadFile(p.abiFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %w", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiContent)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	// Pack arguments for createReport
	reportedAddress := common.HexToAddress(address)
	category := uint8(1) // Example: Phishing
	txData, err := parsedABI.Pack("createReport", reportedAddress, category)
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
	chainID := big.NewInt(p.chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return &entities.PolygonResponse{
		Tsx: signedTx.Hash().Hex(),
	}, nil
}
