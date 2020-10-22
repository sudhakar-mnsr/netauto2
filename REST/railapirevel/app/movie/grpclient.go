package main

import (
	"log"

	pb "github.com/Hands-On-Restful-Web-services-with-Go/chapter6/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
// setup a connection to the server
conn, err := grpc.Dial(address, grpc.WithInsecure())
if err != nil {
   log.Fatal("Did not connect: %v", err)
}

defer conn.Close()
c := pb.NewMoneyTransactionClient(conn)

// Prepare data. Get this from clients like frontend or App
form := "1234"
