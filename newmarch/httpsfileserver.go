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


