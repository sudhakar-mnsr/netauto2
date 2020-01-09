package main

import "fmt"

func main() {
   msg := "Bobsayshelloworld!"
   fmt.Println(msg[:3], msg[7:12], msg[12:17],msg[len(msg)-1:])
   fmt.Println(sort(msg))
   printBytes(msg)
}

