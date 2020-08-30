/* EchoClient
*/
package main

import (
   "fmt"
   "glang.org/x/net/websocket"
   "io"
   "os"
)

func main() {
if len(os.Args) != 2 {
   fmt.Println("Usage: ", os.Args[0], "ws://host:port")
   os.Exit(1)
}
service := os.Args[1]
