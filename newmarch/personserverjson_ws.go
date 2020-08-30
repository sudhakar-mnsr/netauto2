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
   err := websocket.JSON.Receive(ws, &person)
   if err != nil {
      fmt.Println("Cant receive")
   } else {
      fmt.Println("Name: " + person.Name)
      for _, e := range person.Emails {
         fmt.Println("An email: " + e)
      }
   }
}
