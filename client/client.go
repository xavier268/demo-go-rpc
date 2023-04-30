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

	// make a sync call for arith
	a := models.Args{A: 7, B: 8}
	reply := new(models.Quotient)
	// NB : Note that the full service name to use is TYPE.METHOD, not just METHOD !
	err = client.Call("Arith.Divide", &a, reply)
	if err != nil {
		log.Fatal("calling arith :", err)
	}
	fmt.Println(a, " ---> ", *reply)

	// twice !
	err = client.Call("Arith.Divide", &a, nil)
	if err != nil {
		log.Fatal("calling arith :", err)
	}
	fmt.Println(a, " ---> ", *reply)

	// PrintService
	msg := "Hello world !"
	rep := 0
	err = client.Call("PrintService.Print", &msg, &rep)
	if err != nil {
		log.Fatal("calling print:", err)
	} else {
		fmt.Println("Did you see the print ?")
	}

	// make a sync call for arith
	a.A = 50
	// NB : Note that the full service name to use is TYPE.METHOD, not just METHOD !
	err = client.Call("Arith.Divide", &a, reply)
	if err != nil {
		log.Fatal("calling arith :", err)
	}
	fmt.Println(a, " ---> ", *reply)

}
