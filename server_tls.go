package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
   var addr, network, cert, key string
   flag.StringVar(&addr, "e", ":4443", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.StringVar(&cert, "cert", "../certs/localhost-cert.pem", "public cert")
   flag.StringVar(&key, "key", "../certs/localhost-key.pem", "private key")
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
   
   // configure tls with certs and other settings
   tlsConfig := &tls.Config{
      Certificates: []tls.Certificate{cer},
   }
   
   // instead of net.Listen we now use tls.Listen to start a listener
   // on the secure port
   ln, err := nil {
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

