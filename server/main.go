package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	// # starting the server at on the machine
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	// starting the gRPC server on the machines
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the gRPC Server, %v", err)
	}
}
