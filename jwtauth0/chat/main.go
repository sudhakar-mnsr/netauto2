package main

import (
   "log"
   "net/http"
   "text/template"
   "path/filepath"
   "sync"
)

// templ represents a single template
type templateHandler struct {
   once sync.Once
   filename string
   templ *template.Template
}

// 

func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte(`
         <html>
            <head>
               <title>Chat</title>
            </head>
            <body>
               Let's chat!
            </body>
         </html>
      `)) 
   }) 

   // Start the web server
   if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal("ListenAndServe:", err)
   }
}
