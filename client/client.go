package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/xavier268/demo-go-rpc/models"
)

func main() {
	fmt.Println("Start Client")
	client, err := rpc.DialHTTP("tcp", models.Addr)
	if err != nil {
		log.Fatal("dialing:", err)
	} else {
		fmt.Println("Connecting to ", models.Addr)
	}

	// make a sync call
	a := models.Args{A: 7, B: 8}
	reply := new(models.Quotient)
	// NB : Note that the full service name to use is TYPE.METHOD, not just METHOD !
	err = client.Call("Arith.Divide", &a, reply)

	if err != nil {
		log.Fatal("calling :", err)
	}

	fmt.Println(a, " ---> ", *reply)

}
