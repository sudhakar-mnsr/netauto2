/* UDPDaytimeClient
 */
package main

import (
        "fmt"
        "net"
        "os"
)

func main() {
if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
}

service := os.Args[1]

udpAddr, err := net.ResolveUDPAddr("udp", service)
checkError(err)

conn, err := net.DialUDP("udp", nil, udpAddr)
checkError(err)

var buf [512]byte
n, err := conn.Read(buf[0:])

