package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
var addr string
flag.StringVar(&addr, "e", ":4040", "service address endpoint")
flag.Parse()

// create local addr for socket
laddr, err := net.ResolveTCPAddr("tcp", addr)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

// announce services using ListenTCP which is a TCPListener
l, err := net.ListenTCP("tcp", laddr)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}
defer l.Close()
fmt.Println("listening at (tcp)", laddr.String())
