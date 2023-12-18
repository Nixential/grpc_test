package main

import (
	"context"
	pb "grpc_test/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloWorldServiceServer
}

func (s *server) SayHello(context.Context, *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hello World",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error listening on port 5000: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, &server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
