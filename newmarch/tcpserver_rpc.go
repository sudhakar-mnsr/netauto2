/**
* TCPArithServer
 */

package main

import (
	"fmt"
	"net/rpc"
	"errors"
	"net"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
