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

transport := &http.Transport(Proxy: http.ProxyURL(proxyURL))
client := &http.Client(Transport: transport)

request, err := http.NewRequest("GET", url.String(), nil)

urlp, _ := transport.Proxy(request)
fmt.Println("Proxy ", urlp)
dump, _ := httputil.DumpRequest(request, false)
fmt.Println(string(dump))

response, err := client.Do(request)

checkError(err)
fmt.Println("Read ok")

if response.Status != "200 OK" {
   fmt.Println(response.Status)
   os.Exit(2)
}
fmt.Println("Response ok")

var buf [512]byte
reader := response.Body
