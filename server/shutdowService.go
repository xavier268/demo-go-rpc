package main

import (
	"fmt"
	"time"
)

type ShutDownService struct{}

func (*ShutDownService) Close(when *time.Duration, ignore *int) error {
	go func() {

		fmt.Println("killing server in ", when)
		time.Sleep(*when)
		rpcListener.Close()
	}()
	return nil
}
