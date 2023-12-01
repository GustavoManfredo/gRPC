package main

import (
	"client/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{
		Name: "Gustavo",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println(res)
}
