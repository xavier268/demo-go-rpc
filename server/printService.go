package main

import "log"

type PrintService struct{}

func (*PrintService) Print(msg *string, _ *struct{}) error {
	log.Println(*msg)
	return nil
}
