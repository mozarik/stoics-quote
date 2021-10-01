package main

import (
	"context"
	"log"
	"mini-project/hellopb"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	h := hellopb.NewHelloServiceClient(conn)
	response, err := h.SayHello(context.Background(), &hellopb.HelloRequest{Name: "golang"})
	if err != nil {
		log.Fatalf("could not greet: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
