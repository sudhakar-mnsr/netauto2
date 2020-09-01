/**
 * ArithServer
 */

package main

import (
	"fmt"
	"net/rpc"
	"errors"
	"net/http"
)

type Values struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
