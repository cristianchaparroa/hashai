package blockscout

import (
	"fmt"
)

// shortenAddress shortens an Ethereum address to first 6 and last 4 characters
func shortenAddress(address string) string {
	if len(address) > 10 {
		return fmt.Sprintf("%s...%s", address[:6], address[len(address)-4:])
	}
	return address
}

// formatValue formats the value in ETH with appropriate units
func formatValue(value float64) string {
	ethValue := value / 1e18 // Convert from wei to ETH

	if ethValue >= 1 {
		return fmt.Sprintf("%.2f ETH", ethValue)
	} else if ethValue >= 0.001 {
		return fmt.Sprintf("%.2f mETH", ethValue*1000)
	}
	return fmt.Sprintf("%.2f µETH", ethValue*1e6)
}
