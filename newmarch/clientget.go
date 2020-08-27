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
url, err := url.Parse(os.Args[1])
checkError(err)

client := &http.Client{}

request, err := http.NewRequest("HEAD", url.String(), nil)

// only accept UTF-8
request.Header.Add("Accept-Charset", "utf-8;q=1, ISO-8859-1;q=0")
checkError(err)


