package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
var host string
var network string
flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
flag.StringVar(&network, "n", "udp", "network protocol to use")
flag.Parse()

// req data packet is a 48-byte long value
// that is used for sending time request
req := make([]byte, 48)

