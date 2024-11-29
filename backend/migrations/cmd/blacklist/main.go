package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"hashtracker/config"
	"hashtracker/internal/usecases/repo/polygon"
	"hashtracker/pkg/filesys"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	gfs := filesys.NewGitFileSystem()
	filePath, getFileError := gfs.GetGitFilePath("backend/migrations/datasets/addresses-darklist-test.json")
	if getFileError != nil {
		panic(getFileError)
	}

	ethClient, err := ethclient.DialContext(context.Background(), cfg.Polygon.RpcURl)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Polygon node: %v", err))
	}

	contractConfig := &polygon.ContractConfig{
		PrivateKey:      cfg.Polygon.PrivateKey,
		RpcURL:          cfg.Polygon.RpcURl,
		ChainID:         cfg.Polygon.ChainID,
		AbiFilePath:     cfg.BlackListContract.AbiFile,
		ContractAddress: cfg.BlackListContract.Address,
	}

	blacklistRepo, repoErr := polygon.NewBlacklistRepository(contractConfig, ethClient)
	// replace the above line when you want to iterate over the file and not send anything to the blockchain.
	// blacklistRepo, repoErr := polygon.NewBlacklistNoopRepository()
	if repoErr != nil {
		panic(fmt.Sprintf("failed to instantiate the blacklist repository: %v", repoErr))
	}

	l := NewLoader(blacklistRepo)
	loadErr := l.Load(context.Background(), filePath)
	if loadErr != nil {
		panic(loadErr)
	}
}
