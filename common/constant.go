// date: 2019-02-27
package common

//bh.order.subscribe.base.btc_usdt.depth
//bh.order.leverage.btc_usdt.depth
var Order = "bh.order.%s.%s.%s"

//bh.chat.order.history
var Chat = "bh.chat.%s.%s"

//账户订阅区分币币、法币...
//订阅区分全量订阅 与币种订阅
//订阅区分币种
//bh.account.symbol.depth.all
//bh.account.c2c.depth.all
var Account = "bh.account.%s.%s"

var HeartBeat = "bh.heart.ping"

var JWTHeader = "alg"
var JWTKey = "alg"

const (
	Success = 0
	Fail    = 1
)

const (
	MODULE_ORDER     = "order"
	MODULE_ACCOUNT   = "account"
	MODULE_CHAT      = "chat"
	MODULE_HEARTBEAT = "ping"
)

const (
	Subscribe    = "subscribe"
	UnbSubscribe = "un_subscribe"
	Ping         = "ping"
	Pong         = "pong"
)

const (
	ErrorProtocol    = 9001 //协议错误
	ErrorParameter   = 9002 //参数错误
	ErrorDataFalsify = 9003 //数据被篡改
	ErrorSign        = 9004 //签名错误
	ErrorCommand     = 9005 //指令错误
	ErrorHeader      = 9006 //协议头错误
	ErrorData        = 9007 //数据错误
	ErrorVersion     = 9008 //版本错误
	ErrorAuth        = 9009 //认证错误
	ErrorIllegal     = 9010 //非法操作
	ErrorToken       = 9011 //token失效
	ErrorSubRepeat   = 9012 //重复订阅
)
