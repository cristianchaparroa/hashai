package polygon

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BatchReport struct {
	Comments        string
	Source          string
	Category        *big.Int
	Date            *big.Int
	ReportedAddress common.Address
}
