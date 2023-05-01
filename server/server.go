package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/xavier268/demo-go-rpc/models"
)

var rpcListener net.Listener // make it global, so we can close the server ...

// Launch server
func main() {
	var e error

	fmt.Println("Started server")

	rpc.Register(new(ArithService))    // register services
	rpc.Register(new(PrintService))    // register services
	rpc.Register(new(ShutDownService)) // register services
	rpc.Register(new(WatchdogService)) // register services

	rpc.HandleHTTP()                                // register http handler for rpc on the default httpServer
	rpcListener, e = net.Listen("tcp", models.Addr) // create listener
	if e != nil {
		log.Fatal("listen error:", e)
	} else {
		fmt.Println("Listening on ", models.Addr)
	}

	e = http.Serve(rpcListener, nil) // launch default http server, blocking
	fmt.Println("Server stopped : ", e)

}
