package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func main() {
// Set up a /hello resource handler
http.HandleFunc("/hello", helloHandler)

// Create a CA certificate pool and add cert.pem to it
caCert, err := ioutil.ReadFile("cert.pem")
if err != nil {
   log.Fatal(err)
}
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)
