/* EchoServer 
*/

package main

import (
   "fmt"
   "golang.org/x/net/websocket"
   "net/http"
   "os"
)

func Echo(ws *websocket.Conn) {
   fmt.Println("Echoing")
   
   for n := 0; n < 10; n++ {
      msg := "Hello " + string(n+48)
      fmt.Println("Sending to client: " + msg)
      
      err := websocket.Message.Send(ws, msg)
      if err != nil {
         fmt.Println("Cant send")
         break
      } 
      
      var reply string
      err = websocket.Message.Receive(ws, &reply)
      if err != nil {
         fmt.Println("Cant receive")
         break
      } 
      fmt.Println("Received back from client: " + reply)
   }
}
