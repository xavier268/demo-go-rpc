package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/xavier268/demo-go-rpc/models"
)

var listener net.Listener

// Launch server
func main() {
	var e error

	fmt.Println("Started server")

	rpc.Register(new(ArithService))    // register services
	rpc.Register(new(PrintService))    // register services
	rpc.Register(new(ShutDownService)) // register services

	rpc.HandleHTTP()                             // register http handler for rpc on the default httpServer
	listener, e = net.Listen("tcp", models.Addr) // create listener
	if e != nil {
		log.Fatal("listen error:", e)
	} else {
		fmt.Println("Listening on ", models.Addr)
	}

	e = http.Serve(listener, nil) // launch default http server, blocking
	fmt.Println("Server stopped : ", e)

}
