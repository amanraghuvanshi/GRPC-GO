package main

import (
	"log"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

// Stream response send
func (s *helloServer) callSayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got Request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		// adding some delay,
	}
	return nil
}
