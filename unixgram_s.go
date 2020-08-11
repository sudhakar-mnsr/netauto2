package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo server over
// Unix Domain Socket (streaming).  When the server receives a
// request, it returns its content immediately.
func main() {
   var addr string
   flag.StringVar(&addr, "e", "/tmp/echo2.sock", "service endpoint address")
   flag.Parse()
   
   // create local unix domain socket address endpoint
   laddr, err := net.ResolveUnixAddr("unix", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // announce service using ListenUnix which creates a UnixListener
   l, err := net.ListenUnix("unix", laddr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.Close()
   fmt.Println("listening at (unix)", laddr.String())
   
   for {
      // use UnixListener to block and wait for UDS connection request using
      // AcceptUnix which then creates a UnixConn
      conn, err := l.AcceptUnix()
      if err != nil {
         fmt.Println("failed to accept conn:", err)
         conn.Close()
         continue
      }
      fmt.Println("connected to: ", conn.RemoteAddr())
      
      go handleConnection(conn)
   }
}

// handleConnection reads request from connection with conn.Read() then
// write response using conn.Write(). Then the connection is closed.
func handleConnection(conn *net.UnixConn) {
defer conn.Close()
buf := make([]byte, 1024)

n, err := conn.Read(buf)
if err != nil {
   fmt.Println(err)
   return
}
