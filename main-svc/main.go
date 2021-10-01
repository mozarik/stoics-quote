package main

import (
	"fmt"
	"log"
	"mini-project/hellopb"
	"mini-project/main-svc/server"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const ROOT_DIR = "../"

func main() {
	// Set Env
	err := godotenv.Load(ROOT_DIR + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AUTH_API_GATEWAY_PORT := os.Getenv("AUTH_API_GATEWAY_PORT")

	fmt.Println("Go gRPC Server From Main Service at :3001")
	lis, err := net.Listen("tcp", AUTH_API_GATEWAY_PORT)
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
