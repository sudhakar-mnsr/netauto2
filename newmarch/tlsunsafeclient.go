/* TLSUnsafeClientGet
   Go presently bails out when it encounters certificate errors
   However we are configuring to ignore certificate errors
   Transport configuration flag InsecureSkipVerify is set
   (NOT A GOOD PRACTISE)
   With self-signed certificates go will generate error of below
   x509: certificate signed by unknown authority
*/
package main

import (
   "fmt"
   "net/http"
   "net/url"
   "os"
   "strings"
   "crypto/tls"
)

func main() {
if len(os.Args) != 2 {
   fmt.Println("Usage: ", os.Args[0], "https://host:port/page")
   os.Exit(1)
}
url, err := url.Parse(os.Args[1])
checkError(err)

if url.Scheme != "https" {
   fmt.Println("Not https scheme ", url.Scheme)
   os.Exit(1)
}

transport := &http.Transport{}
transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
client := &http.Client{Transport: transport}

request, err := http.NewRequest("GET", url.String(), nil)
// only accept UTF-8
checkError(err)
