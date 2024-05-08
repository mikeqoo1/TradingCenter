package trade_core

import (
	"fmt"
)

type Order struct {
	Product  string
	Type     string // "buy" or "sell"
	Price    float64
	Quantity int
}

// MatchOrders matches buy and sell orders.
func MatchOrders(input chan Order, output chan string) {
	// Placeholder for actual matching logic and report generation
	for order := range input {
		// Placeholder for matching logic
		// Here, just return a sample report
		report := fmt.Sprintf("Order matched for product %s\n", order.Product)
		fmt.Println(report)
		output <- report
	}
}
