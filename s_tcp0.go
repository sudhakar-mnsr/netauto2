package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
   var addr string
   flag.StringVar(&addr, "e", ":4040", "service address endpoint")
   flag.Parse()
   
   // create local addr for socket
   laddr, err := net.ResolveTCPAddr("tcp", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // announce services using ListenTCP which is a TCPListener
   l, err := net.ListenTCP("tcp", laddr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.Close()
   fmt.Println("listening at (tcp)", laddr.String())
   
   for {
      // use TCP listener to block and wait for TCP connection request using
      // AcceptTCP which creates a TCPConn
      conn, err := l.AcceptTCP()
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
// write response using conn.Write(). THen the connection is closed.

func handleConnection(conn *net.TCPConn) {
defer conn.Close()
buf := make([]byte, 1024)
 
n, err := conn.Read(buf)
if err != nil {
   fmt.Println(err)
   return
}

