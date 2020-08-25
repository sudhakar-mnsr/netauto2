package main

import (
   "fmt"
   "net"
   "os"
   "time"
)

func main() {
service := ":1200"
tcpAddr, err := net.ResolveTCPAddr("tcp", service)
checkError(err)

listener, err := net.ListenTCP("tcp", tcpAddr)
checkError(err)

for {
conn, err := listener.Accept()
if err != nil {
   continue
}


