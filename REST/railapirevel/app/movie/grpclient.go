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
