package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"

	. "grpc/client/client"
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
	prodRes, err := prodClient.GetProdStock(context.Background(), &ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}


