/* HTTPSFileServer
   HTTPS = HTTP + TLS
   for a server to use https it need x.509 certificate and a private-
   key file for that certificate. GO at present requires these in -
   PEM-encoded. 
   Then HTTP function ListenAndServe is replaced with the HTTPS -
   HTTP+TLS function ListenAndServeTLS
*/

package main 

import (
   "fmt"
   "net/http"
   "os"
)

func main() {
// deliver files from the directory /var/www
fileServer := http.FileServer(http.Dir("/var/www"))

// register the handler and deliver requests to it
err := http.ListenAndServeTLS(":8000", "jan.mewmarch.name.pem", "private.pem", fileServer)
checkError(err)
}

func checkError(err error) {
   if err != nil {
      fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
