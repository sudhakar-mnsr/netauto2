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
   
   listener, err := net.ListenTCP("tcp", tcpAddr)
   checkError(err)
   
   for {
   conn, err := listener.Accept()
   if err != nil {
      continue
   }
   
   handleClient(conn)
   conn.Close()
   }
}


