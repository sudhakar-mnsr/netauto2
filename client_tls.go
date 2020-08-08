package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	curr "currency/lib"
)

const prompt = "currency"

func main() {
// setup flags
var addr, network, ca string
flag.StringVar(&addr, "e", "localhost:4443", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
flag.StringVar(&ca, "ca", "/tmp/certs/cert.pem", "CA certificate")
flag.Parse()

// Load our CA certificate
caCert, err := ioutil.ReadFile(ca)
if err != nil {
   log.Fatal("failed to read CA cert", err)
}

