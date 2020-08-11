package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
var host string
var network string
flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
flag.StringVar(&network, "n", "udp", "network protocol to use")
flag.Parse()

// req data packet is a 48-byte long value
// that is used for sending time request
req := make([]byte, 48)

// req is initialized with 0x1B or 0001 1011 which is a request setting
// for time server.
req[0] = 0x1B

// rsp byte slice used to receive server response
rsp := make([]byte, 48)

// create a Dialer which allows us to specify dialing options. We will
// need this a bit later to configure the local address when the program
// is using "unixgram"
dialer := net.Dialer{}

// IMPORTANT: when network is "unixgram", the local address must be
// created and set explicitly (see ntpc2.go)
if network == "unixgram" {
   laddr := &net.UnixAddr{Name: fmt.Sprintf("%s-client", host), Net: network}
   dialer.LocalAddr = laddr
}

// setup connection (net.Conn) with Dial()
conn, err := dialer.Dial(network, host)
if err != nil {
   fmt.Printf("failed to connect: %v\n", err)
   os.Exit(1)
}
defer func() {
   if err := conn.Close(); err != nil {
      fmt.Println("failed while closing connection:", err)
   }
}()

fmt.Printf"time from (%s) (%s)\n", network, conn.RemoteAddr())

// send time request
if _, err = conn.Write(req); err != nil {
   fmt.Printf("failed to send request: %v\n", err)
   os.Exit(1)
}

// block to receive server response
read, err := conn.Read(rsp)
if err != nil {
   fmt.Printf("failed to receive response: %v\n", err)
   os.Exit(1)
}
