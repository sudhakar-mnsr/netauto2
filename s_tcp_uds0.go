package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.Parse()
   
   // validate network
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
   	fmt.Println("unsupported network protocol")
   	os.Exit(1)
   }
   
   // anounce service using the listen function which creates a generic
   // listen listener.
   l, err := net.Listen(network, addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.Close()
   fmt.Printf("listening at (%s) %s\n", network, addr)
   
   for {
      // use listener to blcok and wait for connection request using
      // function Accept() which then creates a generic Conn value
      conn, err := l.Accept()
      if err != nil {
         fmt.Println("failed to accept conn:", err)
         conn.Close()
         continue
      }
      fmt.Println("connected to:", conn.RemoteAddr())
      
      go handleConnection(conn)
   }
}

// handleConnection reads request from connection with conn.Read() then
// write response using conn.Write(). Then the connection is closed.
func handleConnection(conn net.Conn) {
   defer conn.Close()
   
   buf := make([]byte, 1024)
   
   n, err := conn.Read(buf)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   w, err := conn.Write(buf[:n])
   if err != nil {
      fmt.Println("failed to write to client:", err)
      return
   }
   if w != n {
      fmt.Println("warning: not all data sent to client")
      return
   }
}
