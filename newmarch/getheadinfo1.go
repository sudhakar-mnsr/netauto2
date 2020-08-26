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

