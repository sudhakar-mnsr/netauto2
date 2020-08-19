package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// This program lists network interface information
func main() {
var ifname string
if len(os.Args) == 2 {
   ifname = os.Args[1]
}
ifaces, err := net.Interfaces()
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

