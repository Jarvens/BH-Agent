// date: 2019-03-01
package grpc

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"os"
)

func BidDirectionalClient() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect fail err: %v\n", err)
		return
	}
	defer conn.Close()
	client := NewRpcPushServiceClient(conn)
	ctx := context.Background()
	stream, err := client.BidStream(ctx)
	if err != nil {
		fmt.Printf("create stream fail err: %v\n", err)
	}
	go func() {
		fmt.Println("请输入消息 ... ")
		input := bufio.NewReader(os.Stdin)
		for {
			line, _ := input.ReadString('\n')
			fmt.Printf("命令行输入: %v\n", line)
			if err := stream.Send(&RpcPushRequest{Event: line, Data: `{"userId":"123456"}`}); err != nil {
				return
			}
		}
	}()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("收到服务端结束信号")
			break
		}
		if err != nil {
			fmt.Println("接收数据出错")
		}
		fmt.Println("客户端收到数据: ", res)
	}
}
