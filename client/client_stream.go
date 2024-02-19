package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

func callsayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started!!")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while sending the names %v", err)
	}

	//  now we will go througn all the names, would frame the response
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the stream %v", err)
		}
		log.Printf("Send the request with values: %s", name)
		time.Sleep(2 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	//  we need to read the resp now
	log.Printf("Client Streaming Finished")

	if err != nil {
		log.Fatalf("Error while recieveing the files, %v", err)
	}
	log.Printf("The message: %v", resp.Message)
}
