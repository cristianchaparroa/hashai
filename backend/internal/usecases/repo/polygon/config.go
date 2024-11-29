package polygon

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
)

type ContractConfig struct {
	PrivateKey      string
	RpcURL          string
	ChainID         int64
	AbiFilePath     string
	ContractAddress string
}

func (c *ContractConfig) GetECDSAPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func (c *ContractConfig) GetContractAddress() common.Address {
	return common.HexToAddress(c.ContractAddress)
}
