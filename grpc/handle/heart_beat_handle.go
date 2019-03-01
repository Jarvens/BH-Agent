// date: 2019-02-28
package handle

import (
	"github.com/Jarvens/BH-Agent/grpc"
)

func HeartBeatHandle(request *grpc.RpcPushRequest, stream grpc.RpcPushService_BidStreamServer) error {
	return nil
}
