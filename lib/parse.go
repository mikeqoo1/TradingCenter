package lib

import (
	"fmt"
	"strconv"
	"strings"
)

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

// ParseMessage 解析電文並且返回電文結構
func ParseEmMessage(msg string) (*EmMessage, error) {
	fmt.Println("ParseMessage=[", msg, "]")
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
