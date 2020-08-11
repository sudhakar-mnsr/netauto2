package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
var addr string
flag.StringVar(&addr, "e", "/tmp/echo2.sock", "service address endpoint")
flag.Parse()
text := flag.Arg(0)
