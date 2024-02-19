package main

import (
	"io"
	"log"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBiDirectionalStreamingServer) error {

	// since both the sides are equally responsible for sending stream and processing them, so we need to process them both

	// first stream that we are going to receive will be of the request
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving the request stream: %v", err)
		}
		log.Printf("Got Request with names: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello, " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while streaming the response: %v", err)
		}
	}
}
