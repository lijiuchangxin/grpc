package service

import (
	"io"
	"log"
	"strconv"
)

type Steamer struct {

}

func (this *Steamer) BidStream(stream Chat_BidStreamServer) error  {
	ctx := stream.Context()
	for {
		select {
		case <- ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错", err)
				return err
			}

			switch res.Input {
			case "对话结束\n":
				log.Println("收到对话结束指令")
				if err := stream.Send(&Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				return nil
			case "返回数据流\n":
				log.Println("收到对话结束指令")
				for i := 0; i < 10; i++{
					if err := stream.Send(&Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}
			default:
				log.Printf("[收到消息]:%s", res.Input)
				if err := stream.Send(&Response{Output: "服务端返回: " + res.Input}); err != nil {
					return err
				}
			}
		}
	}
}