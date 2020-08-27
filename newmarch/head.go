/* HEAD Method
   Every http request type has its own GO function in net/http
   Simplest request from user-agent is HEAD
   This asks information about resource and its HTTP server
   The Status of response is in response field Status
   field Header is map of type header fields in HTTP response
*/

package main

import (
   "fmt"
   "net/http"
   "os"
)

func main() {
if len(os.Args) != 2 {
   fmt.Println("Usage: ", os.Args[0], "host:port")
   os.Exit(1)
}

