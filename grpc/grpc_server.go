// date: 2019-02-28
package grpc

import (
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
	"github.com/Jarvens/BH-Agent/grpc/handle"
	"google.golang.org/grpc"
	"io"
	"net"
	"strings"
)

// rpc连接管理
type RpcConnection struct {
	ChanMap map[string]interface{}
}

var GlobalConnection = new(RpcConnection)

const (
	port = ":3000"
)

type bidServer struct {
}

func (s *bidServer) BidStream(stream RpcPushService_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("receive client message for broken request")
			//TODO 断连处理器
			return ctx.Err()
		default:
			request, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("client stream over")
				return nil
			}
			if err != nil {
				fmt.Printf("receive stream error")
				return err
			}
			module := strings.Split(request.Event, ".")[1]
			switch module {
			case common.MODULE_ORDER:
				fmt.Println("start invoke orderHandle")
				err := handle.OrderHandler(request, stream)
				return err
			case common.MODULE_CHAT:
				fmt.Println("start invoke chatHandle")
				err := handle.ChatHandle(request, stream)
				return err
			case common.MODULE_ACCOUNT:
				fmt.Println("start invoke accountHandle")
				err := handle.AccountHandle(request, stream)
				return err
			case common.MODULE_HEARTBEAT:
				fmt.Println("start invoke heartBeatHandle")
				handle.HeartBeatHandle(request, stream)
			default:
				fmt.Println("command not found")
				if err := stream.Send(&RpcPushResponse{Code: common.ErrorCommand, Message: common.ErrorMsg(common.ErrorCommand), Data: ""}); err != nil {
					return err
				}
			}
		}

	}
}

// gRPC server start
func BidDirectionalServer() {
	server := grpc.NewServer()
	RegisterRpcPushServiceServer(server, &bidServer{})
	address, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}

func (r *RpcConnection) RpcConnectionAdd() error {
	return nil
}
