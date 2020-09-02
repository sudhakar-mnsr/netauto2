/**
* TCPArithClient
 */

package main

import (
	"net/rpc"
	"fmt"
	"log"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
if len(os.Args) != 2 {
   fmt.Println("Usage: ", os.Args[0], "server:port")
   os.Exit(1)
}
service := os.Args[1]

client, err := rpc.Dial("tcp", service)
if err != nil {
   log.Fatal("dialing:", err)
}
