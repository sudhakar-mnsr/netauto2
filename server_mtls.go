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
