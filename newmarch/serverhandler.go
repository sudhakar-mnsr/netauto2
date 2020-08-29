/* ServerHandler
   HTTP requests received by a GO server are usually handled by-
   a multiplexer, which examines the path in the HTTP request and-
   calls appropriate file handler. 
   You can define your handler, this should be registered with -
   default multiplexer by calling http.HandleFunc. The ListenAndServe-
   then take a nil handler.
   If you want to takeover the multiplexer role then you can give-
   ca non-nil function as the handler function to ListenAndServe.
*/

package main

import (
   "net/http"
)

func main() {
   myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
   // just return no content - arbitrary headers can be set, arbitrary body
   rw.WriteHeader(http.StatusNoContent)
   })
   
   http.ListenAndServe(":8080", myHandler)
}  
