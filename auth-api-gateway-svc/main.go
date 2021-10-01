package main

import (
	"context"
	"log"
	"mini-project/hellopb"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const ROOT_DIR = "../"

func main() {
	// Set Config with GoDotEnv\
	err := godotenv.Load(ROOT_DIR + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MAIN_SVC_PORT := os.Getenv("MAIN_SVC_PORT")

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(MAIN_SVC_PORT, grpc.WithInsecure())
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
