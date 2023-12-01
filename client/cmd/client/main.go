package main

import (
	"client/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	reqPerson := &pb.Person{
		Name: "Sure",
	}

	log.Println(res)

	responsePerson, err := client.GetPerson(context.Background(), reqPerson)
	if err != nil {
		log.Fatalf("Error", err)
	}

	log.Println(responsePerson)
}
