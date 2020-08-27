/* ClientGet
*/

package main

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
if len(os.Args) != 2 {
fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
os.Exit(1)
}

