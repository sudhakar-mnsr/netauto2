package main

import (
"crypto/tls"
"crypto/x509"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func main() {
// Read the key pair to create certificate
cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
if err != nil {
   log.Fatal(err)
}
