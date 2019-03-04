// date: 2019-02-28
package grpc

import (
	"encoding/json"
	"github.com/Jarvens/BH-Agent/common"
	"strings"
)

// bh.order.subscribe.base.btc_usdt.web
// 订单处理器
// @param request 请求参数
// @param context 请求上下文
// @return error
func OrderHandler(request *RpcPushRequest, stream RpcPushService_BidStreamServer) (message string, code int) {
	event := strings.Split(request.Event, ".")
	dataMap := make(map[string]string)
	json.Unmarshal([]byte(request.Data), &dataMap)
	userId := dataMap["userId"]
	//token := dataMap["token"]
	//if token == "" {
	//	fmt.Printf("token参数丢失：%s", request.Data)
	//	return "参数错误", common.ErrorParameter
	//}
	//claims, valid := common.ParseToken(token, common.JWTKey)
	//if valid {
	//	tokenUserId := claims.(jwt.MapClaims)["userId"]
	//	if tokenUserId != userId {
	//		return "非法操作", common.ErrorIllegal
	//	}
	//} else {
	//	return "token失效", common.ErrorToken
	//}

	// module 模块 order account chat
	// eventType 事件类型  subscribe ub_subscribe
	// moduleType 模块类型 base leverage
	// symbol 交易对  all ... btc_usdt
	// clientType 客户端类型  web  pc
	module := event[1]
	eventType := event[2]
	moduleType := event[3]
	symbol := event[4]
	clientType := event[5]
	if eventType == common.Subscribe {
		subscribe(request.Event, module, moduleType, symbol, clientType, userId, stream)
	} else if eventType == common.UnbSubscribe {
		//unSubscribe(module, symbol, subTypes, userId, stream)
	}
	return "", common.Success
}

// 订阅
// module 模块 order account chat
// eventType 事件类型  subscribe ub_subscribe
// moduleType 模块类型 base leverage
// symbol 交易对  all ... btc_usdt
// clientType 客户端类型  web  pc
func subscribe(evt, module, moduleType, symbol, clientType,
	userId string, stream RpcPushService_BidStreamServer) (message string, code int32) {

	if moduleType == "base" {
	} else if moduleType == "leverage" {
	}
	valid, _ := common.Contain(userId, GlobalConnection)
	//存在连接
	if valid {
		userMap := GlobalConnection[userId]
		chanMap := userMap.(map[string]interface{})
		connMap := chanMap[clientType]
		moduleMap := connMap.(map[interface{}]interface{})
		eventMap := moduleMap[stream]
		event := eventMap.(map[string]interface{})
		value := event[module]
		val := value.([]string)
		//判断订阅信息是否存在 避免重复订阅
		valid, _ := common.Contain(evt, value)
		if valid {
			return "重复订阅", common.ErrorSubRepeat
		} else {
			val = append(val, evt)
			return "订阅成功", common.Success
		}

	} else {
		event := []string{evt}
		moduleMap := make(map[string]interface{})
		moduleMap[common.MODULE_ORDER] = event

		connMap := make(map[interface{}]interface{})
		connMap[stream] = moduleMap

		chanType := make(map[string]interface{})
		chanType["web"] = connMap

		userMap := make(map[string]interface{})
		userMap[userId] = chanType
		GlobalConnection = userMap
		return "订阅成功", common.Success
	}
	return "订阅失败", common.Success
}

// 取消订阅
// @param module 模块 币币、杠杆
// @param symbol 交易对 btc_usdt all
// @param unSubType 订阅类型 增量  全量
// @param userId 用户id
func unSubscribe(module, symbol, unSubType, userId string, stream RpcPushService_BidStreamServer) {

}
