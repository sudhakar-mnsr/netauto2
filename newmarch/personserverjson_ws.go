/* PersonServerJSON
 */
package main

import (
   "fmt"
   "golang.org/x/net/websocket"
   "net/http"
   "os"
)

type Person struct {
   Name string
   Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
var person Person
