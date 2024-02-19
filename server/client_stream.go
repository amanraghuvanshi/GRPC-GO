package main

import (
	"io"
	"log"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string // slice of messages, since here the client is going to send the stream, and the server needs to process it

	for {
		req, err := stream.Recv()
		// processing the stream
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}

		if err != nil {
			return err
		}
		log.Printf("Got Request with names: %v", req.Name)
		messages = append(messages, "Hello, ", req.Name)
	}
}
