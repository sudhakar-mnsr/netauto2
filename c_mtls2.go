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

// Create a CA certificate pool and add cert.pem to it
caCert, err := ioutil.ReadFile("cert.pem")
if err != nil {
   log.Fatal(err)
}

caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

// Create a HTTPS client and supply the created CA pool and certificate
client := &http.Client{
          Transport: &http.Transport{
             TLSClientConfig: &tls.Config{
                RootCAs: caCertPool,
                Certificates: []tls.Certificate{cert},
             },
          },
}
