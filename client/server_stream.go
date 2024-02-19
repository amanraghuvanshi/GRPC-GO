package main

import (
	"context"
	"io"
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

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming the data %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming Finished!")

}
