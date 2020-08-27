/* GET
   Retrieve representation of a resource rather than just information
   Content of response is in response field Body (type io.ReadCloser).
 */

package main

import (
        "fmt"
        "net/http"
        "net/http/httputil"
        "os"
        "strings"
)

func main() {
if len(os.Args) != 2 {
   fmt.Println("Usage: ", os.Args[0], "host:port")
   os.Exit(1)
}
url := os.Args[1]

response, err := http.Get(url)
if err != nil {
   fmt.Println(err.Error())
   os.Exit(2)
}

if response.Status != "200 OK" {
   fmt.Println(response.Status)
   os.Exit(2)
}
