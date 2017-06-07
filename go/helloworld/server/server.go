package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "../../grpc-stub/generated"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	newServer := grpc.NewServer()
	pb.RegisterGreeterServer(newServer, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(newServer)
	if err := newServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}