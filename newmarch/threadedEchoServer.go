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

func handleClient(conn net.Conn) {
   defer conn.Close()
   
   var buf [512]byte
   for {
      // read upto 512 bytes
      n, err := conn.Read(buf[0:])
      if err 	= nil {
         return
      }
      fmt.Println(string[buf[0:]))
      // Write the n bytes read
      _, err2 := conn.Write(buf[0:n])
      if err2 != nil {
         return
      }
   }
}

