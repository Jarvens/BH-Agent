// date: 2019-02-28
package handle

import (
	"github.com/Jarvens/BH-Agent/grpc"
)

//资产处理器
// 订单处理器
// @param request 请求参数
// @param context 请求上下文
// @return error
func AccountHandle(request *grpc.RpcPushRequest, stream grpc.RpcPushService_BidStreamServer) error {
	return nil
}