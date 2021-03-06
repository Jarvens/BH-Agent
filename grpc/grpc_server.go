// date: 2019-02-28
package grpc

import (
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
	"go/constant"
	"google.golang.org/grpc"
	"io"
	"net"
	"strings"
)

var GlobalConnection = make(map[string]interface{})

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
				fmt.Println("关闭客户端")
				return nil
			}
			if err != nil {
				fmt.Printf("接收数据流错误: %v\n", err)
				return err
			}

			str := strings.Split(request.Event, ".")
			if len(str) <= 1 {
				streamSend(stream, "指令错误", int32(common.ErrorCommand))
			} else {
				module := str[1]
				switch module {
				case common.MODULE_ORDER:
					message, code := OrderHandler(request, stream)
					streamSend(stream, message, int32(code))
				case common.MODULE_CHAT:
					err := ChatHandle(request, stream)
					return err
				case common.MODULE_ACCOUNT:
					err := AccountHandle(request, stream)
					return err
				case common.MODULE_HEARTBEAT:
					HeartBeatHandle(request, stream)
				default:
					fmt.Printf("指令不存在: %v\n", module)
					streamSend(stream, "指令不存在", int32(constant.Unknown))
				}
			}
		}

	}
}

//启动gRPC服务
func BidDirectionalServer() {
	//1.实例化gRPC
	//2.注册服务
	//3.监听端口
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

//发送流数据
func streamSend(stream RpcPushService_BidStreamServer, message string, code int32) {
	err := stream.Send(&RpcPushResponse{Message: message, Code: code})
	if err != nil {
		fmt.Printf("Stream send message has error: %s,%d \n", message, code)
	}
}
