package main

import (
	"context"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Server Responds with: Hello",
	}, nil
}
