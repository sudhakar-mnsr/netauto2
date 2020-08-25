package main 

import (
   "fmt"
   "net"
   "os"
)

func main() {
service := ":1201"
tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
checkError(err)

