package main

import (
	"github.com/xavier268/demo-go-rpc/models"
)

// Define Arith service
type Arith struct{}

func (t *Arith) Divide(args *models.Args, quo *models.Quotient) error {
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
