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

// use ResolveUnixAddr to create remote address to server
raddr, err := net.ResolveUnixAddr("unix", addr)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

// Use DialUnix to create a connection to the remote address.
// Note: there is no requirement to specify the local address
conn, err := net.DialUnix("unix", nil, raddr)
if err != nil {
   fmt.Println("failed to connect to server:", err)
   os.Exit(1)
}
defer conn.Close()

// send text to server
_, err = conn.Write([]byte(text))
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

