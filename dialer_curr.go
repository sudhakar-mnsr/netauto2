package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"learning-go/ch11/curr1"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

const prompt = "currency"

func main() {
// setup flags
var addr string
var network string
flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
flag.Parse()
