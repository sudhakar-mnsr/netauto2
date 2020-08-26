/* Ping
 Prepare an IP connection
 Prepare ICMP echo command
 Send a ping request to host
 Get a reply
*/
package main

import (
        "bytes"
        "fmt"
        "io"
        "net"
        "os"
)

// set the IP address of your choise
const myIPAddress = "192.168.1.2"
const ipv4HeaderSize = 20
