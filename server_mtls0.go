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

func main() {
   http.HandleFunc("/hello", helloHandler)
    
   // Create a CA certificate pool and add cert.pem to it
   caCert, err := ioutil.ReadFile("cert.pem")
   if err != nil {
      log.Fatal(err)
   }
   
   caCertPool := x509.NewCertPool()
   caCertPool.AppendCertsFromPEM(caCert)
   
   // Cerate the TLS Config with the CA pool and enable CLient certificate
   // validation
   tlsConfig := &tls.Config{
                 ClientCAs: caCertPool,
                 ClientAuth: tls.RequireAndVerifyCLientCert,
   }
   tlsConfig.BuildNameToCertificate()
   
   // Create a server instance to listen on port 8443 with the TLS config
   server := &http.Server{
              Addr: ":8443",
              TLSConfig: tlsConfig,
   }
   
   // Listen to HTTPS connections with the server certificate and wait
   log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
