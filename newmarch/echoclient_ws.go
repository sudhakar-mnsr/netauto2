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

conn, err := websocket.Dial(service, "", "http://localhost:12345")
checkError(err)
var msg string
for {
err := websocket.Message
if err != nil {
   if err == io.EOF {
      break
   }
   fmt.Println("Couldnt receive msg " + err.Error())
   break
}
