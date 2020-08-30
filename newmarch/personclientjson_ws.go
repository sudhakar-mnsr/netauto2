/* PersonClientJSON
 */
package main

import (
   "fmt"
   "golang.org/x/net/websocket"
   "os"
)

type Person struct {
   Name   string
   Emails []string
}

func main() {
   if len(os.Args) != 2 {
           fmt.Println("Usage: ", os.Args[0], "ws://host:port")
           os.Exit(1)
   }
   service := os.Args[1]


