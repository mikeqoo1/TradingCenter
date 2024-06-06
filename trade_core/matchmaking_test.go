package trade_core

import "testing"

func TestMatchmaking_MatchOrders(t *testing.T) {
	type fields struct {
		BuyOrders  []Order
		SellOrders []Order
	}
	type args struct {
		input  chan Order
		output chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "測試案例",
			fields: fields{
				BuyOrders:  []Order{{Product: "1234", Type: "buy", Price: 5.5, Quantity: 1000}},
				SellOrders: []Order{{Product: "1234", Type: "sell", Price: 5, Quantity: 1000}},
			},
			args: args{
				input:  make(chan Order, 1),
				output: make(chan string, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matchmaking{
				BuyOrders:  tt.fields.BuyOrders,
				SellOrders: tt.fields.SellOrders,
			}
			go m.MatchOrders(tt.args.input, tt.args.output)

			// 模擬輸入
			tt.args.input <- tt.fields.BuyOrders[0]
			close(tt.args.input)

			// 檢查輸出
			result := <-tt.args.output
			if result != "撮合成功" {
				t.Errorf("預期 '撮合成功', 但是得到 '%s'", result)
			}
		})
	}
}
