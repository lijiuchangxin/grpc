package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/server/service"
	"io/ioutil"
	"net/http"
)

func main() {

	//creds, _ := credentials.NewClientTLSFromFile("keys/server.crt", "keys/server.key")
	//rpcServer := grpc.NewServer(grpc.Creds(creds))

	//rpcServer := grpc.NewServer()

	//加载密钥对
	cert, _ := tls.LoadX509KeyPair("server/cert/server.pem", "server/cert/server.key")
	// 创建证书池
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates:	[]tls.Certificate{cert},  			// 服务端证书
		ClientAuth: 	tls.RequireAndVerifyClientCert,		// 用于双向验证
		ClientCAs:		certPool,							// 指定客户端证书池

	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))


	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))
	service.RegisterOderServiceServer(rpcServer, new(service.OrderService))
	service.RegisterChatServer(rpcServer, new(service.Steamer))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	//_ = httpServer.ListenAndServeTLS("keys/server.crt", "keys/server.key")

	a := httpServer.ListenAndServeTLS("server/cert/server.pem", "server/cert/server.key")
	fmt.Println(a)
}


// test
