package polygon

import (
	"context"
	"hashtracker/config"
	"hashtracker/internal/entities"
	"hashtracker/internal/usecases"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


type polygonRepository struct {
	privateKey string
	rpcURL string
	chainID int64
	abiFilePath string
	contractAddress string
}

func NewPolygonRepository(abiFilePath string, contractAddress string, chainID int64, cfg config.Config) usecases.PolygonRepository {
	return &polygonRepository{
		privateKey: cfg.PolygonPrivateKey,
		rpcURL: cfg.PolygonRpcURl,
		chainID: chainID,
		abiFilePath: abiFilePath,
		contractAddress: contractAddress,
	}
}

func (p *polygonRepository) Resolve (ctx context.Context) (*entities.PolygonResponse, error) {
	// Connect to Polygon Mumbai
	client, err := ethclient.Dial(p.rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Polygon node: %v", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(p.privateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(p.chainID)) // 80002 Polygon Amoy
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Contract address
	contractAddress := common.HexToAddress(p.contractAddress)

	// Load ABI from a file
	abiContent, err := os.ReadFile(p.abiFilePath)
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiContent)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// Pack arguments for createReport
	reportedAddress := common.HexToAddress("0x0")
	category := uint8(1) // Example: Phishing
	txData, err := parsedABI.Pack("createReport", reportedAddress, category)
	if err != nil {
		log.Fatalf("Failed to pack arguments for createReport: %v", err)
	}

	// Get current nonce
	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to fetch nonce: %v", err)
	}

	// Estimate gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &contractAddress,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     txData,
	})
	if err != nil {
		log.Fatalf("Failed to estimate gas: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch gas price: %v", err)
	}

	// Construct transaction
	tx := types.NewTransaction(
		nonce,          // Nonce
		contractAddress, // Contract address
		big.NewInt(0),   // Value (sending 0 MATIC in this case)
		gasLimit,        // Gas limit
		gasPrice,        // Gas price
		txData,          // Encoded data
	)

	// Sign transaction
	chainID := big.NewInt(p.chainID) // Polygon Mumbai chain ID
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	return &entities.PolygonResponse{
		Tsx: signedTx.Hash().Hex(),
	}, nil
}
