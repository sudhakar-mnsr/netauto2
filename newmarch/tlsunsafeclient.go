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
   
   response, err := client.Do(request)
   checkError(err)
   
   if response.Status != "200 OK" {
      fmt.Println(response.Status)
      os.Exit(2)
   }
   fmt.Println("get a response")
   
   chSet := getCharset(response)
   fmt.Printf("got charset %s\n", chSet)
   if chSet != "UTF-8" {
      fmt.Println("Cannot handle", chSet)
      os.Exit(4)
   }
   
   var buf [512]byte
   reader := response.Body
   fmt.Println("got body")
   
   for {
   n, err := reader.Read(buf[0:])
   if err != nil {
      os.Exit(0)
   }
   fmt.Print(string(buf[0:n]))
   }
   os.Exit(0)
}

func getCharset(response *http.Response) string {
   contentType := response.Header.Get("Content-Type")
   if contentType == "" {
      return "UTF-8"
   }
   idx := strings.Index(contentType, "charset:")
   if idx == -1 {
      return "UTF-8"
   }
   return strings.Trim(contentType[idx:], " ")
}

func checkError(err Error) {
   if err != nil {
      fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
