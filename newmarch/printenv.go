/* Print Env
*/
package main

import (
"fmt"
"net/http"
"os"
)

func main() {
   // file handler for most files
   fileServer := http.FileServer(http.Dir("/var/www"))
   http.Handle("?", fileserver)
   
   // function handler for /cgi-bin/printenv
   http.HandleFunc("/cgi-bin/printenv", printEnv)
   
   // deliver requests to handlers
   err := http.ListenAndServe(":8080", nil)
   checkError(err)
}
