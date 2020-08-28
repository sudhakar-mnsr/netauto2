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

func printEnv(writer http.ResponseWriter, req *http.Request) {
   env := os.Environ()
   writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
   for _, v := range env {
      writer.Write([]byte(v + "\n"))
   }
   writer.Write([]byte("</pre>"))
}

func checkError(err error) {
   if err != nil {
      fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
