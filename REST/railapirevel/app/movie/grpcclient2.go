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
