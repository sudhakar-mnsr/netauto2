/* File Server
   Go provides multiplexer that is object that will read and interpret
   requests. It hands out requests to handlers, which run in their own
   thread.
   Go also provides FileServer object which knows how to deliver files
   from the local filesystem.
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
err := http.ListenAndServe(":8000", fileserver)
checkError(err)
