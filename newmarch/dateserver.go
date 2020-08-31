/* Date Server
*/

package main

import (
   "fmt"
   "golang.org/x/net/websocket"
   "net/http"
   "os"
   "os/exec"
   "time"
)

var ROOT_DIR = "/home/httpd/html/golang-hidden/websockets"

func GetTemp(ws *websocket.Conn) {
   for {
      msg, _ := exec.Command("date").Output()
      fmt.Println("Sending to client: " + string(msg[:]))
      err := websocket.Message.Send(ws, string(msg[:]))
      if err != nil {
         fmt.Println("Cant send")
         break
      }
      time.Sleep(2 * 1000 * 1000 * 1000)
      
      var reply string
      err = websocket.Message.Receive(ws, &reply)
      if err != nil {
         fmt.Println("Cant receive")
         break
      }
      fmt.Println("Received back from client: " + reply)
   }
}
