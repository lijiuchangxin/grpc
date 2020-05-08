package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	. "grpc/server/service"
)

func main() {
	//creds, _ := credentials.NewClientTLSFromFile("keys/server.crt", "demo.joyios.com")

	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	// 创建证书池
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates:	[]tls.Certificate{cert},  			// 客户端证书
		ServerName: 	"localhost",
		RootCAs:		certPool,
	})


	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds) )

	defer func() {
		conn.Close()
	}()

	if err != nil {
		log.Fatal(err)
	}

	//prodClient := NewProdServiceClient(conn)
	prodClient := NewProdServiceClient(conn)
	orderClient := NewOderServiceClient(conn)
	chatClient := NewChatClient(conn)
	ctx := context.Background()

	t := timestamp.Timestamp{Seconds:time.Now().Unix()}
	res, _ := orderClient.NewOrder(ctx, &OrderMain{
		OrderId:    1001,
		OrderNo:    "20190809",
		UserId:     1,
		OrderMoney: 112,
		OrderTime:	&t,
	})
	fmt.Println(res)


	prodRes, err := prodClient.GetProdStock(ctx, &ProdRequest{ProdId: 12, ProdArea:ProdAreas_C})  // 获取商品库存
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)

	response, err := prodClient.GetProdStocks(ctx, &QuerySize{Size: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Prodres)


	info, err := prodClient.GetProdInfo(ctx, &ProdRequest{ProdId: 1, ProdArea:ProdAreas_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

	stream, err := chatClient.BidStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("请输入消息...")
		input := bufio.NewReader(os.Stdin)
		for {
			input_, _ := input.ReadString('\n')
			if err := stream.Send(&Request{Input: input_}); err != nil {
				return
			}
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("⚠️ 收到服务端的结束信号")
		}

		if err != nil {
			log.Println("接受数据出错", err)
		}

		log.Printf("[客户端接收到]: %s", resp.Output)
	}
}


