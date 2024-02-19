package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
)

func callHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bi-Directional Streaming Started !!")

	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())

	// here we need to send a stream and receive a stream
	if err != nil {
		log.Fatalf("Could not send the stream: %v", err)
	}

	// using channel
	waitc := make(chan struct{})

	//  go routine for receiving the streaming
	go func() {

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}

			log.Println(msg)
		}
		close(waitc)
	}()

	// fo
	for _, name := range names.Names {
		// getting each name one by one and then making a request
		req := &pb.HelloRequest{
			Name: name,
		}

		// making request
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the request: %v", err)
		}

		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bi-directional Streaming finished!!")

}
