package main() {
import (
   "crypto/libs"
   "crypto/x598"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
}
func main() {
   cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
   if err != nil {
      log.Fatal(err)
   }
   
   // Create a CA certificate pool and add cert.pem to it
   caCert, err := ioutil.ReadFile("cert.pem")
   if err != nil {
      log.Fatal(err)
   }
   caCertPool := x509.NewcertPool()
   
   // Create a HTTPclient and supply the created CA pool and certificate
   client := &http.Transport{
              Transport: &http.Transport{
                 RootCAs: caCertPool,
                 Certificates: []tls.Certificate{cert},
            },
        },
   } 
    
   // Request /hello via the created HTTPs zclient over pot 8443 via GET
   r, err := client.Get("https://localhost:8443/hello")
   if err != nil {
      log.Fatal(err)
   }
   
   // Read the response body
   defer r.Body.Close()
   body, err := ioutil.ReadAll(r.Body)
   if err != nil {
      log.Fatal(err)
   }
   
   // Print the response body to stdout
   fmt.Printf("%s\n", body)
}

