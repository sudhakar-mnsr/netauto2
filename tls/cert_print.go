package main

import (
   "crypto/tls"
   "fmt"
   "certinfo"
   "log"
)

func main() {
   // Connect to google.com
   cfg := tls.Config{}
   conn, err := tls.Dial("tcp", "google.com:443", &cfg)
   if err != nil {
      log.Fatalln("TLS connection failed: " + err.Error())
   }
   
   // Grab the last certificate in the chain
   certChain := conn.ConnectionState.PeerCertificates
   cert := certChain[len(certChain)-1]
   
   // Print the certificate
   result, err := certinfo.CertificateText(cert)
   if err != nil {
      log.Fatal(err)
   }
   
   fmt.Print(result)
}
