package server

import (
	"context"
	"log"
	"mini-project/hellopb"
)

type Server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Printf("Receive message body from client: %s", req.Name)
	return &hellopb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
