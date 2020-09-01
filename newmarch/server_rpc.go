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

type Arith int

func (t *Arith) Multiply(args *Values, reply *int) error {
   *reply = args.A * args.B
   return nil
}

func (t *Arith) Divide(args *Values, quo *Quotient) error {
   if args.B == 0 {
      return errors.New("divide by zero")
   }
   quo.Quo = args.A / args.B
   quo.Rem = args.A % args.B
   return nil
}
