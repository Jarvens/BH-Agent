// date: 2019-02-28
package grpc

import (
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
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

			str := strings.Split(request.Event, ".")
			if len(str) <= 1 {
				fmt.Printf("event content error: %s\n", request.Event)
				if err := stream.Send(&RpcPushResponse{Code: common.ErrorProtocol, Message: common.ErrorMsg(common.ErrorProtocol), Data: ""}); err != nil {
					return err
				}
			} else {
				module := str[1]
				switch module {
				case common.MODULE_ORDER:
					fmt.Println("start invoke orderHandle")
					message, code := OrderHandler(request, stream)
					if code != common.Success {

					}
					return err
				case common.MODULE_CHAT:
					fmt.Println("start invoke chatHandle")
					err := ChatHandle(request, stream)
					return err
				case common.MODULE_ACCOUNT:
					fmt.Println("start invoke accountHandle")
					err := AccountHandle(request, stream)
					return err
				case common.MODULE_HEARTBEAT:
					fmt.Println("start invoke heartBeatHandle")
					HeartBeatHandle(request, stream)
				default:
					fmt.Println("command not found")
					if err := stream.Send(&RpcPushResponse{Code: common.ErrorCommand, Message: common.ErrorMsg(common.ErrorCommand), Data: ""}); err != nil {
						return err
					}
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

func StreamSend(stream RpcPushService_BidStreamServer, message string, code int32) {
	stream.Send(&RpcPushResponse{Message: message, Code: code})
}
