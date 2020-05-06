package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/service"
	"net"
)

func main() {

	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	listen, err := net.Listen("tcp", ":8081")
	fmt.Println(err)

	rpcServer.Serve(listen)

	//listener, err := net.Listen("tcp", "8081")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//log.Printf("listen on: %s\n", "8801")
	//
	//server := grpc.NewServer()
	//
	//service.RegisterProdServiceServer(server, &service.ProdService{})
	//
	//
	//if err := server.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
