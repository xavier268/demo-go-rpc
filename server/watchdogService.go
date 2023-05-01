package main

import (
	"fmt"
	"sync"
	"time"
)

type WatchdogService struct{}

var ticker *time.Ticker
var tickerGuard sync.Mutex

var defaultDuration = time.Second * 10

func init() {
	fmt.Println("Setting initial ticker at ", defaultDuration)
	ticker = time.NewTicker(defaultDuration)
	go func() {
		for range ticker.C {
			{
				fmt.Print(".")
				// fmt.Println("Ticker received ", t.UnixMilli()%100_000)
				// do something, like shuting down, or sending a ping, or checking health, ...
			}
		}
	}()

}

// Arm will set up the watchdog. If no ping received during a duration intervall, triggers event ..
func (wd *WatchdogService) Arm(dur *time.Duration, _ *struct{}) error {
	fmt.Println("Changing ticker to ", *dur)
	tickerGuard.Lock()
	ticker.Reset(*dur)
	tickerGuard.Unlock()
	return nil
}
