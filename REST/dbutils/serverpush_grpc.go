package main

import (
   "fmt"
   "log"
   "net"
   "time"
   pb "github.com/narenaryan/serverPush/datafiles"
   "google.golang.org/grpc"
   "google.golang.org/grpc/reflection"
)

const (
   port = ":50051"
   noOfSteps = 3
)

// server is used to create MoneyTransactionServer.
type server struct{}
