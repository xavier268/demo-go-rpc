// Shared package between server and clients
package models

var Addr = ":1234" // service adress

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}
