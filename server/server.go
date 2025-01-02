package server

import (
	"fmt"
	"net"
	"strings"

	"TradingCenter/lib"
	"TradingCenter/trade_core"
)

// Server struct holds information about the TCP server.
type Server struct {
	Address string
}

// NewServer creates a new TCP server with the given address.
func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

// Start starts the TCP server.
func (s *Server) Start(orderChannel chan trade_core.Order, reportChannel chan string) error {
	// Listen on the specified address
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("error listening: %v", err)
	}
	defer listener.Close()
	fmt.Printf("Server listening on %s\n", s.Address)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		fmt.Println("Client connected")

		// Handle client requests in a separate goroutine
		go s.handleClient(conn, orderChannel, reportChannel)
	}
}

// handleClient handles client connections.
func (s *Server) handleClient(conn net.Conn, orderChannel chan trade_core.Order, reportChannel chan string) {
	defer conn.Close()

	// Read client request
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	// Parse Msg
	/*
		上市櫃:
			委託
			委託8=FIX.4.4^A9=0222^A35=D^A34=00000136^A49=T845JJb^A56=XTAI^A52=20241227-02:12:23.516^A50=8450^A57=0^A1=9713322^A11=0D0598BC0002^A37=30002^A54=1^A55=1234  ^A60=20241227-10:13:23.086^A38=000001^A40=2^A59=0^A44=00030.0000^A10000=1^A10001=0^A10002=0^A1080= ^A10004=N^A10=211^A
			回報
			回報8=FIX.4.4^A9=0251^A35=8^A49=XTAI^A56=T845JJb^A34=136^A52=20241227-02:13:23.091^A50=0^A57=8450^A37=30002^A11=0D0598BC0002^A17=0D0598BC0002^A150=0^A39=0^A1=9713322^A38=1^A40=2^A59=0^A44=30.0000^A55=1234  ^A54=1^A60=20241227-02:13:23.090^A32=0^A151=1^A14=0^A6=0^A31=0.0000^A10000=1^A10001=0^A10002=0^A10=209^A
		興櫃:
			委託
			新單(興櫃)韻玲姐給交易所的委託8=FIX.4.3^A9=0223^A35=UO01^A34=00000102^A49=845T3131^A52=20240506-07:30:21^A56=emgMsgSvr^A1=9826947^A11=00001^A38=00001000^A44=00098.0000^A54=1^A55=1269  ^A76=8450^A97=N^A117=00001^A80001=03^A80002=03^A80003=01^A80004=0000^A80014=2BK00001^A80024=153021^A81001=1^A10=118^A
			改單(興櫃)韻玲姐給交易所的委託8=FIX.4.3^A9=0186^A35=UO02^A34=00000434^A49=845T3131^A52=20240506-08:44:21^A56=emgMsgSvr^A80001=03^A80002=03^A80003=02^A80004=0000^A80014=2BK00055^A80024=164421^A81013=0000059^A11=0004j^A38=00000000^A44=00026.0000^A97=N^A10=062^A
			刪單(興櫃)韻玲姐給交易所的委託8=FIX.4.3^A9=0160^A35=UO03^A34=00000502^A49=845T3131^A52=20240506-08:58:37^A56=emgMsgSvr^A80001=03^A80002=03^A80003=03^A80004=0000^A80014=2BK0005I^A80024=165837^A81013=0000077^A11=0004u^A97=N^A10=182^A
			回報
			委託(興櫃)交易所給韻玲姐的回報8=FIX.4.3^A9=296^A35=UO20^A49=emgMsgSvr^A56=845T3131^A34=1205^A52=20240506-07:30:20^A80001=03^A80002=03^A80003=20^A80014=2BK00001^A80024=153021^A80004=0000^A11=00001^A81010=153021931^A81060=0000001^A81061=00000.0000^A81062=00001000^A81063=0000001^A81064=00098.0000^A81065=00001000^A76=8450^A117=00001^A1=9826947^A55=1269  ^A81001=1^A54=1^A10=202^A
			成交(興櫃)交易所給韻玲姐的回報8=FIX.4.3^A9=268^A35=UC24^A49=emgMsgSvr^A56=845T3131^A34=2716^A52=20240506-08:22:10^A80001=03^A80002=05^A80003=24^A80014=        ^A80024=162210^A80004=0000^A81011=1^A375=815T^A81003=EMTS_PVC3 ^A81013=0000001^A1=9826947^A55=1269  ^A54=1^A44=00098.0000^A38=00001000^A81002=162211408^A81026=0001850^A17=0000003^A10=029^A
	*/
	//fmt.Println("bytes=[", buffer, "]")
	fmt.Println("buffet=[", string(buffer), "]")

	// Unmarshal JSON into order struct
	var order trade_core.Order

	if strings.Contains(string(buffer), "XTAI") || strings.Contains(string(buffer), "ROCO") {
		fmt.Println("上市櫃委託")
		msg, err := lib.ParseTwseMessage(string(buffer))
		if err != nil {
			fmt.Println(err.Error())
		}
		order.Exchange = 1
		order.Product = msg.MsgStockID
		if msg.MsgBScode == 1 {
			order.Type = "buy"
		} else {
			order.Type = "sell"
		}
		order.Price = msg.MsgPrice
		order.Quantity = msg.MsgQty
		order.Twse_Allmsg = *msg
	} else {
		fmt.Println("興櫃委託")
		msg, err := lib.ParseEmMessage(string(buffer))
		if err != nil {
			fmt.Println(err.Error())
		}
		order.Exchange = 2
		order.Product = msg.MsgStockID
		if msg.MsgBScode == 1 {
			order.Type = "buy"
		} else {
			order.Type = "sell"
		}
		order.Price = msg.MsgPrice
		order.Quantity = msg.MsgQty
		order.Em_Allmsg = *msg
	}

	// Send order to orderChannel for processing
	order.SendOrderReport = false
	orderChannel <- order

	// Listen for reports from reportChannel
	go func() {
		for {
			report := <-reportChannel
			_, err := conn.Write([]byte(report))
			if err != nil {
				fmt.Printf("Error writing: %v\n", err)
				return
			}
		}
	}()
}
