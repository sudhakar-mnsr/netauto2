package main

import (
	"log"
	"net"

	pb "github.com/Hands-On-Restful-Web-services-with-Go/chapter6/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to create MoneyTransactionServer.
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTrancaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
   log.Printf("Got request for money transfer ....")
   log.Printf("Amount: %f, from A/c: %s", in.Amount, in.From, in.To)
   // Do database logic here....
   return &pb.TransactionResponse(Confirmation: true}, nil
}
