package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/xavier268/demo-go-rpc/models"
)

// Launch server
func main() {
	fmt.Println("Started server")

	rpc.Register(new(Arith))        // register services
	rpc.Register(new(PrintService)) // register services

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
