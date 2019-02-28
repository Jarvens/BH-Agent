// date: 2019-02-28
package grpc

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)
const (
	port=":3000"
)

type server struct {
	
}

func (s *server)BidStream(stream RpcPushService_BidStreamServer)error  {
	ctx:=stream.Context()
	for{
		select{
		case <-ctx.Done():
			fmt.Printf("收到客户端通过context发出的关闭信号")
			//TODO 断连处理器
			return ctx.Err()
		default:
			data, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				fmt.Printf("接收数据错误")
				return err
			}
			//bh.order.base.btc_usdt.depth
			module:=strings.Split(data.Event,".")[1]
			switch module {
			case "结束对话\n":
				fmt.Printf("收到结束对话指令")
				if err := stream.Send(&ChatResponse{Output: "收到结束指令"}); err != nil {
					return err
				}
				return nil
			case "返回数据流\n":
				fmt.Printf("收到返回数据流指令")
				for i := 0; i < 10; i++ {
					if err := stream.Send(&ChatResponse{Output: "服务端返回：" + data.Input}); err != nil {
						return err
					}
				}
			default:
				b,_:=json.Marshal(ctx)
				fmt.Printf("打印上下文信息: %v\n",string(b))
				fmt.Printf("收到消息：%s\n", data.Input)
				if err := stream.Send(&ChatResponse{Output: "服务器端返回：" + data.Input}); err != nil {
					return err
				}

			}
		}

		}
	}
}
