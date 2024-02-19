package main

import (
	"context"
	"log"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

// server streaming function
func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming Started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Couldnt send the names: %v", err)
	}

}
