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

// create a dialer to configure its settings instead of using default
// dialer from net.Dial() function
dialer := &net.Dialer{
   Timeout: time.Second * 300,
   KeepAlive: time.Minute * 5,
}

// simple dialing strategy with retry with a simple backoff.
// More sophisticated retry strategies follow similar pattern but may
// include features such as exponential backoff delay, etc
var (
   conn net.conn
   err error
   connTries = 0
   connMaxRetries = 3
   connSleepRetry = time.Second * 1
)

for connTries < connMaxRetries {
fmt.Println("creating connection socket to", addr)
conn, err = dialer.Dial(network, addr)
if err != nil {
   fmt.Println("failed to create socket:", err)
   switch nerr := err.(type) {
   case net.Error:
      if nerr.Temporary() {
         connTries++
         fmt.Println("trying again in:", connSleepRetry)
         time.Sleep(connSleepRetry)
         continue
      }
      fmt.Println("unable to recover")
      os.Exit(1)
   default:
      os.Exit(1)
   }
}
