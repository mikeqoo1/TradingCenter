package trade_core

import (
	"TradingCenter/lib"
	"fmt"
)

type Order struct {
	Product         string
	Type            string // "buy" or "sell"
	Price           float64
	Quantity        int
	Exchange        int // 1表示上市櫃 2表示興櫃 3表示期貨
	SendOrderReport bool
	Twse_Allmsg     lib.TwseMessage
	Em_Allmsg       lib.EmMessage
}

type Matchmaking struct {
	BuyOrders  []Order
	SellOrders []Order
}

func NewMatchmaking() *Matchmaking {
	return &Matchmaking{
		BuyOrders:  make([]Order, 0),
		SellOrders: make([]Order, 0),
	}
}

// MatchOrders matches buy and sell orders.
func (m *Matchmaking) MatchOrders(input chan Order, output chan string) {
	for order := range input {
		// 先送出委託回報
		if !order.SendOrderReport {
			order_report := ""
			if order.Exchange == 1 {
				order_report = lib.CreatTwseOrderReport(&order.Twse_Allmsg)
			} else if order.Exchange == 2 {
				order_report = lib.CreatEmOrderReport(&order.Em_Allmsg)
			}
			order.SendOrderReport = true
			fmt.Println(order_report)
			output <- order_report
		}

		if order.Type == "buy" {
			m.BuyOrders = append(m.BuyOrders, order)
		} else if order.Type == "sell" {
			m.SellOrders = append(m.SellOrders, order)
		}

		// Perform matching
		for _, buyOrder := range m.BuyOrders {
			for _, sellOrder := range m.SellOrders {
				if buyOrder.Product == sellOrder.Product && buyOrder.Price >= sellOrder.Price {
					// Match found
					quantity := buyOrder.Quantity
					if sellOrder.Quantity < quantity {
						quantity = sellOrder.Quantity
					}

					// Handle matched orders (this is a placeholder, you may need more sophisticated logic)
					// report := fmt.Sprintf("Matched: %+v with %+v, Quantity: %d\n", buyOrder, sellOrder, quantity)
					if order.Exchange == 1 {
						// 組出成交回報
					} else if order.Exchange == 2 {
						// 組出成交回報
					}
					exec_report := "撮合成功"
					fmt.Println(exec_report)
					output <- exec_report

					// Update quantities
					buyOrder.Quantity -= quantity
					sellOrder.Quantity -= quantity

					// Remove orders with quantity 0
					if buyOrder.Quantity == 0 {
						m.BuyOrders = removeOrder(m.BuyOrders, buyOrder)
					}
					if sellOrder.Quantity == 0 {
						m.SellOrders = removeOrder(m.SellOrders, sellOrder)
					}
				}
			}
		}
	}
}

func removeOrder(orders []Order, order Order) []Order {
	for i, o := range orders {
		if o == order {
			orders[i] = orders[len(orders)-1]
			return orders[:len(orders)-1]
		}
	}
	return orders
}
