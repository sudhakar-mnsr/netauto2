/* Ping
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
