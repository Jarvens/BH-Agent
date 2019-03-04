// date: 2019-02-28
package grpc

import (
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
)

//心跳
func HeartBeatHandle(request *RpcPushRequest, stream RpcPushService_BidStreamServer) {
	if err := stream.Send(&RpcPushResponse{Message: common.Pong, Code: common.Success}); err != nil {
		fmt.Printf("心跳错误: %v\n", err)
	}
}
