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
)

type Loader interface {
	Load(ctx context.Context, filePath string) error
}

type loader struct {
	polygonRepository usecases.PolygonRepository
}

func NewLoader(pr usecases.PolygonRepository) Loader {
	return &loader{
		polygonRepository: pr,
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

	for decoder.More() {
		var address ReportedAddress
		decodeErr := decoder.Decode(&address)
		if decodeErr != nil {
			nErrors++
			fmt.Printf("error decoding address entry: %v\n", decodeErr)
			continue
		}
		count++
		tx, createErr := l.polygonRepository.CreateReport(ctx, &polygon.ReportRequest{
			Address:  address.Address,
			Comments: address.Comment,
			Source:   "hashai-migration",
			Date:     big.NewInt(time.Now().Unix()),
		})
		if createErr != nil {
			fmt.Printf("error creating report entry on Polygon: %v\n", createErr)
			continue
		}
		// Print progress every 1000 entries
		if count%1000 == 0 {
			fmt.Printf("Processed %d addresses...\n", count)
		}
		fmt.Printf("https://amoy.polygonscan.com/tx/%s\n", tx.HashTransaction)
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
