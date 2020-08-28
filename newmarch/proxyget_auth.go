/* ProxyAuthGet
*/

package main

import (
   "encoding/base64"
   "fmt"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

const auth = "jannewmarch:mypassword"

func main() {
if len(os.Args) != 3 {
   fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
   os.Exit(1)
}
proxy := os.Args[1]
proxyURL, err := url.Parse(proxy)
checkError(err)
rawURL := os.Args[2]
url, err := url.Parse(rawURL)
checkError(err)
