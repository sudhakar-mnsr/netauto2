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
   io.WriteString(w, "Hello, world!\n")
}
