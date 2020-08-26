/* IPGetHeadInfo ver1
*/

package main

import (
   "bytes"
   "fmt"
   "io"
   "net"
   "os"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
   }
   service := os.Args[1]
   
   conn, err := net.Dial("tcp", service)
   checkError(err)

   _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
   checkError(err)

   result, err := readFully(conn)
   checkError(err)
   
   fmt.Println(string(result))
   
   os.Exit(0)
}
