/* ThreadedEchoServer
 */
package main

import (
        "fmt"
        "net"
        "os"
)

func main() {
   service := ":1201"
   tcpAddr, err := net.ResolveTCPAddr("tcp", service)
   checkError(err)
   
   listener, err := net.ListenTCP("tcp", tcpAddr)
   checkError(err)
   
   for {
      conn, err := listener.Accept()
      if err != nil {
         continue
      }
      
      // run as goroutine
      go handleClient(conn)
   }
}
