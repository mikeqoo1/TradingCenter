package main

import (
	"fmt"

	"TradingCenter/server"
	"TradingCenter/trade_core"
)

func main() {
	// Create channels for orders and reports
	orderChannel := make(chan trade_core.Order)
	reportChannel := make(chan string)

	// Start order matching process
	go trade_core.MatchOrders(orderChannel, reportChannel)

	// Create and start the server
	serverAddr := "localhost:8080"
	tcpServer := server.NewServer(serverAddr)

	fmt.Printf("Starting server on %s\n", serverAddr)
	go func() {
		err := tcpServer.Start(orderChannel, reportChannel)
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	// Wait indefinitely
	select {}
}
