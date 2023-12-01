package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/pb"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Request:", in)
	return &pb.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *Server) GetPerson(ctx context.Context, in *pb.Person) (*pb.PersonResponse, error) {
	log.Println("Reqeust:", in)
	return &pb.PersonResponse{
		Id: 5,
		Name: in.GetName(),
	}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again, " + in.GetName()}, nil
}

func main() {
	log.Printf("Started running gRPC server")

	// Abrindo um listener
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// Registrando o servidor
	pb.RegisterGreeterServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
