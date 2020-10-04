package main

import (
  "io"
  "log"

  pb "github.com/narenaryan/serverPush/datafiles"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  address = "localhost:50051"
)

// RecieveStream listens to the stream contents and use them
func RecieveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
log.Println("Started listening to the server stream!")
stream, err := client.MakeTransaction(context.Background(), request)
if err != nil {
   log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
}
