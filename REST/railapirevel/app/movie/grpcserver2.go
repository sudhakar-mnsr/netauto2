package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Hands-On-Restful-Web-services-with-Go/chapter6/serverPush/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

// server is used to create MoneyTransactionServer.
type server struct{}
