package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// TwseMessage 上市櫃的格式
type TwseMessage struct {
	FixVersion         string // Tag:8
	MsgLength          int    // Tag:9
	MsgType            string // Tag:35
	MsgSeqNum          int    // Tag:34 應用訊息的訊息序號，由1開始編列。
	MsgSendID          string // Tag:49 傳送端代號=845T3131
	MsgSendTime        string // Tag:52 訊息傳送時間 (YYYYMMDD-HH:MM:SS.sss)
	MsgRecvID          string // Tag:56 接收端代號，集中：XTAI及櫃檯：ROCO
	MsgOrigSendingTime string // Tag:122 原始訊息傳送時間 (YYYYMMDD-HH:MM:SS.sss)
	/*以上都是檔頭*/

	MsgAccountNo       int     // Tag:1 帳號
	MsgTicketNumber    string  // Tag:11 電文序號
	MsgOrderNumber     string  // Tag:37 委託單號
	MsgQty             int     // Tag:38 數量
	MsgMarketOrderType int     // Tag:40 市價或限價類型 (1: 市價, 2: 限價)
	MsgPrice           float64 // Tag:44 價格
	MsgBrokerID        string  // Tag:50 證券商代碼
	MsgBScode          int     // Tag:54 買賣別：“1” - 買，“2” - 賣
	MsgStockID         string  // Tag:55 股票代號
	MsgTimeInForce     int     // Tag:59 委託有效期間
	MsgOrdertime       string  // Tag:60 訊息時間
	MsgTargetSubID     string  // Tag:57 TargetSubID：交易盤別(1碼)，一般交易為0、盤後零股交易為2、盤後定價交易為7、盤中零股為C、標借為4、拍賣為5、一般標購為6、證金標購為B
	MsgOrderPlatform   int     // Tag:10000 委託管道
	MsgCreditType      int     // Tag:10001 信用類別
	MsgOrderKind       int     // Tag:10002 委託類型 (一般、定價:0, 零股:2)
	MsgRefOrderID      int     // Tag:1080 參考訂單 ID
}

// ParseTwseMessage 解析電文並且返回電文結構
func ParseTwseMessage(msg string) (*TwseMessage, error) {
	fmt.Println("ParseTwseMessage=[", msg, "]")
	parts := strings.Split(msg, "\x01")
	fmt.Println("parts=[", parts, "]")
	message := &TwseMessage{}

	for _, part := range parts {
		tagValue := strings.Split(part, "=")
		fmt.Println("tagValue=[", tagValue, "]")
		if len(tagValue) == 2 {
			tag := tagValue[0]
			fmt.Println("tag=[", tag, "]")
			value := tagValue[1]
			fmt.Println("value=[", value, "]")

			switch tag {
			case "8":
				message.FixVersion = value
			case "9":
				lennss, _ := strconv.Atoi(value)
				message.MsgLength = lennss
			case "35":
				message.MsgType = value
			case "34":
				seqnum, _ := strconv.Atoi(value)
				message.MsgSeqNum = seqnum
			case "49":
				message.MsgSendID = value
			case "52":
				message.MsgSendTime = value
			case "56":
				message.MsgRecvID = value
			case "122":
				message.MsgOrigSendingTime = value
			case "1":
				acc, _ := strconv.Atoi(value)
				message.MsgAccountNo = acc
			case "11":
				message.MsgTicketNumber = value
			case "37":
				message.MsgOrderNumber = value
			case "38":
				qqqtttyyy, _ := strconv.Atoi(value)
				message.MsgQty = qqqtttyyy
			case "40":
				marketOrderType, _ := strconv.Atoi(value)
				message.MsgMarketOrderType = marketOrderType
			case "44":
				priceee, _ := strconv.ParseFloat(value, 64)
				message.MsgPrice = priceee
			case "50":
				message.MsgBrokerID = value
			case "54":
				bs, _ := strconv.Atoi(value)
				message.MsgBScode = bs
			case "55":
				message.MsgStockID = value
			case "59":
				timeInForce, _ := strconv.Atoi(value)
				message.MsgTimeInForce = timeInForce
			case "60":
				message.MsgOrdertime = value
			case "57":
				message.MsgTargetSubID = value
			case "10000":
				orderPlatform, _ := strconv.Atoi(value)
				message.MsgOrderPlatform = orderPlatform
			case "10001":
				creditType, _ := strconv.Atoi(value)
				message.MsgCreditType = creditType
			case "10002":
				orderKind, _ := strconv.Atoi(value)
				message.MsgOrderKind = orderKind
			case "1080":
				refOrderID, _ := strconv.Atoi(value)
				message.MsgRefOrderID = refOrderID
			default:
				fmt.Println("Unknown tag:", tag)
			}
		} else {
			break
		}
	}

	return message, nil
}

// CreatTwseOrderReport 上市櫃交易所委託回報
func CreatTwseOrderReport(msg *TwseMessage) string {
	delimiter := "\x01"
	//report := "8=FIX.4.4^A9=0251^A35=8^A49=XTAI^A56=T845JJb^A34=136^A52=20241227-02:13:23.091^A50=0^A57=8450^A37=30002^A11=0D0598BC0002^A17=0D0598BC0002^A150=0^A39=0^A1=9713322^A38=1^A40=2^A59=0^A44=30.0000^A55=1234  ^A54=1^A60=20241227-02:13:23.090^A32=0^A151=1^A14=0^A6=0^A31=0.0000^A10000=1^A10001=0^A10002=0^A10=209^A"
	report := "8=" + msg.FixVersion + delimiter + "150=0" + delimiter
	return report
}

// EmMessage 興櫃的格式
type EmMessage struct {
	FixVersion  string // Tag:8
	MsgLength   int    // Tag:9
	MsgType     string // Tag:35
	MsgSeqNum   int    // Tag:34 應用訊息的訊息序號，由1開始編列。
	MsgSendID   string // Tag:49 傳送端代號=845T3131
	MsgSendTime string // Tag:52 訊息傳送時間 (YYYYMMDD-HH:MM:SS)
	MsgRecvID   string // Tag:56 接收端代號=emgMsgSvr

	/*以上都是檔頭*/

	MsgAccountNo    int     // Tag:1 帳號
	MsgTicketNumber string  // Tag:11 電文序號
	MsgQty          int     // Tag:38 數量
	MsgPrice        float64 // Tag:44 價格
	MsgBScode       int     // Tag:54 買賣別：“1” - 買，“2” - 賣
	MsgStockID      string  // Tag:55 股票代號
	MsgBrokerID     string  // Tag:76 證券商代碼
	MsgResend       string  // Tag:97 重送旗號 "N"-New,"Y"-Resend
	MsgOrderNumber  string  // Tag:117 委託單號
	MsgSystemType   int     // Tag:80001 系統別
	MsgFunctionID   int     // Tag:80002 系統別
	MsgStatusCode   int     // Tag:80004 狀態碼
	MsgUserDefined  int     // Tag:80014 使用者自訂
	MsgTime         string  // Tag:80024 訊息時間 HHMMSS
	MsgOrderKind    int     // Tag:81001 委託類別:1-張,台兩 2-股
	MsgChecksum     int     // Tag:10 檢查碼
}

// ParseEmMessage 解析電文並且返回電文結構
func ParseEmMessage(msg string) (*EmMessage, error) {
	fmt.Println("ParseEmMessage=[", msg, "]")
	parts := strings.Split(msg, "\x01")
	fmt.Println("parts=[", parts, "]")
	message := &EmMessage{}

	for _, part := range parts {
		tagValue := strings.Split(part, "=")
		fmt.Println("tagValue=[", tagValue, "]")
		if len(tagValue) == 2 {
			tag := tagValue[0]
			fmt.Println("tag=[", tag, "]")
			value := tagValue[1]
			fmt.Println("value=[", value, "]")

			switch tag {
			case "8":
				message.FixVersion = value
			case "9":
				lennss, _ := strconv.Atoi(value)
				message.MsgLength = lennss
			case "35":
				message.MsgType = value
			case "34":
				seqnum, _ := strconv.Atoi(value)
				message.MsgSeqNum = seqnum
			case "49":
				message.MsgSendID = value
			case "52":
				message.MsgSendTime = value
			case "56":
				message.MsgRecvID = value
			case "1":
				acc, _ := strconv.Atoi(value)
				message.MsgAccountNo = acc
			case "11":
				message.MsgTicketNumber = value
			case "38":
				qqqtttyyy, _ := strconv.Atoi(value)
				message.MsgQty = qqqtttyyy
			case "44":
				priceee, _ := strconv.ParseFloat(value, 64)
				message.MsgPrice = priceee
			case "54":
				bs, _ := strconv.Atoi(value)
				message.MsgBScode = bs
			case "55":
				message.MsgStockID = value
			case "76":
				message.MsgBrokerID = value
			case "97":
				message.MsgResend = value
			case "117":
				message.MsgOrderNumber = value
			case "80001":
				systype, _ := strconv.Atoi(value)
				message.MsgSystemType = systype
			case "80002":
				funcid, _ := strconv.Atoi(value)
				message.MsgFunctionID = funcid
			case "80004":
				code, _ := strconv.Atoi(value)
				message.MsgStatusCode = code
			case "80014":
				userdefined, _ := strconv.Atoi(value)
				message.MsgUserDefined = userdefined
			case "80024":
				message.MsgTime = value
			case "81001":
				kind, _ := strconv.Atoi(value)
				message.MsgOrderKind = kind
			case "10":
				check, _ := strconv.Atoi(value)
				message.MsgChecksum = check
			default:
				fmt.Println("Unknown tag:", tag)
			}
		} else {
			break
		}
	}

	return message, nil
}

// CreatEmOrderReport 興櫃交易所委託回報
func CreatEmOrderReport(msg *EmMessage) string {
	delimiter := "\x01"
	//report := "8=FIX.4.3^A9=296^A35=UO20^A49=emgMsgSvr^A56=845T3131^A34=1205^A52=20240506-07:30:20^A80001=03^A80002=03^A80003=20^A80014=2BK00001^A80024=153021^A80004=0000^A11=00001^A81010=153021931^A81060=0000001^A81061=00000.0000^A81062=00001000^A81063=0000001^A81064=00098.0000^A81065=00001000^A76=8450^A117=00001^A1=9826947^A55=1269  ^A81001=1^A54=1^A10=202^A"
	report := "8=" + msg.FixVersion + delimiter + "35=UO20" + delimiter
	return report
}
