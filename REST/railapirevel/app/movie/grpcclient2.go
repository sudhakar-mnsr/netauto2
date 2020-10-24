package main

import (
	"io"
	"log"

	pb "github.com/Hands-On-Restful-Web-services-with-Go/chapter6/serverPush/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}

	// Listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// If there are no more messages, get out of loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}
