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
