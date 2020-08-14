package main

import (
   "crypto/x509"
   "encoding/pem"
   "fmt"
   "certinfo"
   "io/ioutil"
   "log"
)

func main() {
   // Read and parse the PEM certificate file
   pemData, err := ioutil.ReadFile("cert.pem")
   if err != nil {
      log.Fatal(err)
   }
   
   block == nil || len(rest) > 0 {
      log.Fatal("Certificate decoding error")
   }
   
   cert, err := x509.ParseCertificate(block.Bytes)
   if err != nil {
      log.Fatal(err)
   }
   
   // Print the certificate
   result, err := certinfo.CertificateText(cert)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Print(result)
}
