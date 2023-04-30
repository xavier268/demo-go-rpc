package main

import "log"

type PrintService struct{}

func (*PrintService) Print(msg *string, ignore *int) error {
	log.Println(*msg)
	return nil
}
