// date: 2019-02-28
package handle

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Jarvens/BH-Agent/common"
	"github.com/Jarvens/BH-Agent/grpc"
	"strings"
)

//bh.order.subscribe.base.btc_usdt.depth
// 订单处理器
// @param request 请求参数
// @param context 请求上下文
// @return error
func OrderHandler(request *grpc.RpcPushRequest, context context.Context) error {
	event := strings.Split(request.Event, ".")
	dataMap := make(map[string]string)
	json.Unmarshal([]byte(request.Data), &dataMap)
	userId := dataMap["userId"]
	if "" == userId {
		return errors.New("参数错误")
	}
	eventType := event[2]
	module := event[3]
	symbol := event[4]
	subTypes := event[5]
	if eventType == common.Subscribe {
		subscribe(module, symbol, subTypes, userId)
	} else if eventType == common.UnbSubscribe {
		unSubscribe(module, symbol, subTypes, userId)
	}
	return nil
}

// 订阅
// @param module 模块 币币、杠杆
// @param symbol 交易对 btc_usdt all
// @param subType 订阅类型 增量  全量
// @param userId  用户id
func subscribe(module, symbol, subType, userId string) {
	if module == "base" {

	} else if module == "leverage" {

	}

}

// 取消订阅
// @param module 模块 币币、杠杆
// @param symbol 交易对 btc_usdt all
// @param unSubType 订阅类型 增量  全量
// @param userId 用户id
func unSubscribe(module, symbol, unSubType, userId string) {

}
