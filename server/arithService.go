package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/xavier268/demo-go-rpc/models"
)

// Define Arith service
type Arith struct{}

func (t *Arith) Divide(args *models.Args, quo *models.Quotient) error {
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// Launch server
func main() {
	fmt.Println("Started server")
	a := new(Arith)                        // create new service
	rpc.Register(a)                        // register service
	rpc.HandleHTTP()                       // register http handler for rpc on the default httpServer
	l, e := net.Listen("tcp", models.Addr) // create listener
	if e != nil {
		log.Fatal("listen error:", e)
	} else {
		fmt.Println("Listening on ", models.Addr)
	}
	e = http.Serve(l, nil) // launch default http server, blocking
	fmt.Println("Server stopped : ", e)

}
