package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

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
	err = client.Call("ArithService.Divide", &a, reply)
	if err != nil {
		log.Fatal("calling arith :", err)
	}
	fmt.Println(a, " ---> ", *reply)

	// twice !
	err = client.Call("ArithService.Divide", &a, reply)
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
	err = client.Call("ArithService.Divide", &a, reply)
	if err != nil {
		log.Fatal("calling arith :", err)
	}
	fmt.Println(a, " ---> ", *reply)

	// kill server
	when := time.Millisecond * 2000
	rep = 0
	err = client.Call("ShutDownService.Close", &when, &rep)
	if err != nil {
		log.Fatal("calling close", err)
	}

}
