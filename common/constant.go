// date: 2019-02-27
package common

//bh.order.base.btc_usdt.depth
//bh.order.leverage.btc_usdt.depth
var Order = "bh.order.%s.%s.%s"

//bh.chat.order.history
var Chat = "bh.chat.%s.%s"

//bh.account.btc_usdt.depth
var Account = "bh.account.%s.%s"

const (
	Success = 0
	Fail    = 1
)

const (
	ChatOrder = "order"
)

const (
	Subscribe   = "subscribe"
	UbSubscribe = "un_subscribe"
	Ping        = "ping"
	Pong        = "pong"
)

const (
	ErrorProtocol    = 9001
	ErrorParameter   = 9002
	ErrorDataFalsify = 9003
	ErrorSign        = 9004
	ErrorCommand     = 9005
	ErrorHeader      = 9006
	ErrorData        = 9007
	ErrorVersion     = 9008
	ErrorAuth        = 9009
)

func ErrorMsg(code int) string {
	switch code {
	case ErrorProtocol:
		return "协议错误"
	case ErrorParameter:
		return "参数错误"
	case ErrorDataFalsify:
		return "数据被篡改"
	case ErrorSign:
		return "签名错误"
	case ErrorCommand:
		return "指令不存在"
	case ErrorHeader:
		return "协议头错误"
	case ErrorData:
		return "数据错误"
	case ErrorVersion:
		return "协议版本错误"
	case ErrorAuth:
		return "认证错误"
	default:
		return "UNKNOWN"
	}
}