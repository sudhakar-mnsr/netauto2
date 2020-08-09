package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../data.csv")
)

func main() {
   // setup flags
   var addr, network, cert, key, ca string
   flag.StringVar(&addr, "e", ":4443", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.StringVar(&cert, "cert", "/tmp/certs/cert.pem", "public cert")
   flag.StringVar(&key, "key", "/tmp/certs/key.pem", "private key")
   flag.StringVar(&ca, "ca", "/tmp/certs/cert.pem", "root CA certificate")
   flag.Parse()
   
   // validate supported network protocols
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
   	fmt.Println("unsupported network protocol")
   	os.Exit(1)
   }
   
   // load server cert by providing the private key that generated it
   cer, err := tls.LoadX509KeyPair(cert, key)
   if err != nil {
      log.Fatal(err)
   }
   
   // load root CA
   caCert, err := ioutil.ReadFile(ca)
   if err != nil {
      log.Fatal(err)
   }
   caPool := x509.NewCertPool()
   caPool.AppendCertsFromPEM(caCert)
   
   // configure tls with certs and other settings
   tlsConfig := &tls.Config{
                 ClientAuth: tls.RequireAndVerifyClientCert,
                 ClientCAs: caPool,
                 Certificates: []tls.Certificate{cer},
   }
   
   // instead of net.Listen, we now use tls.Listen to start a listener 
   // on the secure port
   ln, err := tls.Listen(network, addr, tlsConfig)
   if err != nil {
      log.Println(err)
   }
   defer ln.Close()
   log.Println("***** Global Currency Service (secure) *****")
   log.Printf("Service started: (%s) %s; server cert %s\n", network, addr, cert)
   
   acceptDelay := time.Millisecond * 10
   acceptCount := 0
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         switch e := err.(type) {
         case net.Error:
            // if temporary error, attempt to connect again
            if e.Temporary() {
               if acceptCount > 5 {
                  log.Fatalf("unable to connect after %d retries: %v", acceptCount, err)
               }
               acceptDelay *= 2
               acceptCount++
               time.Sleep(acceptDelay)
               continue
            }
         default:
            log.Println(err)
            if err := conn.Close(); err != nil {
               log.Fatal(err)
            }
            continue
         }
         acceptDelay = time.Millisecond * 10
         acceptCount = 0
      }
      log.Println("securely connected to remote client ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error closing connection:", err)
      }
   }()
   
   // set initial deadline prior to entering the client request/response
   // loop to 45 seconds. This means that the client has 45 seconds to 
   // send its initial request or loose the connection.
   if err := conn.SetDeadline(time.Now().Add(time.Second * 45)); err != nil {
      log.Println("failed to set deadline:", err)
      return
   }
   
   for {
      dec := json.NewDecoder(conn)
      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            if err.Timeout() {
               log.Println("deadline reached, disconnecting...")
            }
            log.Println("network error:", err)
            return
         default:
            if err == io.EOF {
               log.Println("closing connection:", err)
               return
            }
            enc := json.NewEncoder(conn)
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               log.Println("failed error encoding", encerr)
               return
            }
            continue
         }
      }
      // search currencies, result is []curr.Currency
      result := curr.Find(currencies, req.Get)
      
      // send result
      enc := json.NewEncoder(conn)
      if err := enc.Encode(&result); err != nil {
         switch err := err.(type) {
         case net.Error:
            log.Println("failed to send response:", err)
            return
         default:
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               log.Println("failed to send error:", encerr)
               return
            }
            continue
         }
      }
      
      if err := conn.SetDeadline(time.Now().Add(time.Second * 90)); err != nil {
         log.Println("failed to set deadline:", err)
         return
      }
   }
}
