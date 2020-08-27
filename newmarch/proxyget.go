/* ProxyGet
*/

package main

import (
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
if len(os.Args) != 3 {
   fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
   os.Exit(1)
}

proxyString := os.Args[1]
proxyURL, err := url.Parse(proxyString)
checkError(err)
rawURL := os.Args[2]
url, err := url.Parse(rawURL)
checkError(err)
