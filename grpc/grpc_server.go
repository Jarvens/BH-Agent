// date: 2019-02-28
package grpc

import (
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
	"github.com/Jarvens/BH-Agent/handle"
	"io"
	"strings"
)

const (
	port = ":3000"
)

type server struct {
}

func (s *server) BidStream(stream RpcPushService_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("收到客户端通过context发出的关闭信号")
			//TODO 断连处理器
			return ctx.Err()
		default:
			request, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				fmt.Printf("接收数据错误")
				return err
			}
			//bh.order.base.btc_usdt.depth
			module := strings.Split(request.Event, ".")[1]
			switch module {
			case common.MODULE_ORDER:
				fmt.Println("start invoke orderHandle")
				err := handle.OrderHandler(request, ctx)
				return err
			case common.MODULE_CHAT:
				fmt.Println("start invoke chatHandle")
				err := handle.ChatHandle(request, ctx)
				return err
			case common.MODULE_ACCOUNT:
				fmt.Println("start invoke accountHandle")
				err := handle.AccountHandle(request, ctx)
				return err
			case common.MODULE_HEARTBEAT:
				fmt.Println("start invoke heartBeatHandle")
				handle.HeartBeatHandle(request, ctx)
			default:
				fmt.Println("command not found")
				if err := stream.Send(&RpcPushResponse{Code: common.ErrorCommand, Message: common.ErrorMsg(common.ErrorCommand), Data: ""}); err != nil {
					return err
				}
			}
		}

	}
}
