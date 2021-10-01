package main

import (
	"fmt"
	"log"
	"mini-project/hellopb"
	"mini-project/main-svc/server"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Server From Main Service at :3001")
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := server.Server{}
	grpcServer := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
