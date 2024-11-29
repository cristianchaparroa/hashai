package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hashtracker/internal/entities/polygon"
	"hashtracker/internal/usecases"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Loader interface {
	Load(ctx context.Context, filePath string) error
}

type loader struct {
	blacklistRepo usecases.BlacklistReportRepository
}

func NewLoader(br usecases.BlacklistReportRepository) Loader {
	return &loader{
		blacklistRepo: br,
	}
}

func (l *loader) Load(ctx context.Context, filePath string) error {
	file, openErr := os.Open(filePath)
	if openErr != nil {
		return fmt.Errorf("error opening file: %v", openErr)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	_, tokenErr := decoder.Token()
	if tokenErr != nil {
		return fmt.Errorf("error reading opening bracket: %v", tokenErr)
	}

	count := 0
	nErrors := 0
	startTime := time.Now()
	reports := make([]*polygon.BatchReport, 0)

	for decoder.More() {
		var address ReportedAddress
		decodeErr := decoder.Decode(&address)
		if decodeErr != nil {
			nErrors++
			fmt.Printf("error decoding address entry: %v\n", decodeErr)
			continue
		}
		count++
		r := &polygon.BatchReport{
			ReportedAddress: common.HexToAddress(address.Address),
			Comments:        address.Comment,
			Source:          "hashai-migration",
			Date:            big.NewInt(time.Now().Unix()),
		}

		reports = append(reports, r)
	}

	createReportErr := l.blacklistRepo.CreateBatchReport(ctx, reports)
	if createReportErr != nil {
		return fmt.Errorf("error creating batch report: %v", createReportErr)
	}

	// Print final statistics
	duration := time.Since(startTime)
	fmt.Printf("\nProcessing complete:\n")
	fmt.Printf("Total addresses processed: %d\n", count)
	fmt.Printf("Total addresses with errors: %d\n", nErrors)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("Average speed: %.2f addresses/second\n", float64(count)/duration.Seconds())

	return nil
}
