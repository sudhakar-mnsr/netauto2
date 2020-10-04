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
   // Listen to the stream of messages
   for {
      response, err := stream.Recv()
      if err == io.EOF {
         // if there are not more messages, get out of loop
      break
      }
      if err != nil {
         log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
      }
      log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
   }
}

func main() {
   // Setup a connection to the server.
   conn, err := grpc.Dial(address, grpc.WithInsecure())
   if err != nil {
      log.Fatal("Did not connect: %v", err)
   }
   defer conn.Close()
   client := pb.NewMoneyTransactionClient(conn)
   
   // Prepare data. Get this from clients like Front-end or Android App
   from := "1234"
   to := "5678"
   amount := float32(1250.75)
   
   // Contact the server and print out its response.
   RecieveStream(client, &pb.TransactionRequest(From: from, To: to, Amount: amount})
}
