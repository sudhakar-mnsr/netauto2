package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
var addr string
flag.StringVar(&addr, "e", "localhost:4040", "service address endpoint")
flag.Parse()
text := flag.Arg(0)

// use ResolveTCPAddr to create address to connect to
raddr, err := net.ResolveTCPAddr("tcp", addr)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

