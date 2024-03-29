package main

import (
	"log"

	pb "github.com/amanraghuvanshi/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// setting up the connection
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect %v", err)
	}

	// ensuring runs at the very end
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Aman", "Alice", "Bob"},
	}

	// for the unary API, uncomment
	// callSayHello(client)

	// callsayHelloClient()
	//server-side streaming
	// callSayHelloServerStream(client, names)

	//client-size streaming
	// callsayHelloClientStream(client, names)

	// bi-directional streaming
	callHelloBiDirectionalStreaming(client, names)

}
