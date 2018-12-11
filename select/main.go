package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Channel 1 message"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Channel 2 message"
	}()

	var result string

	select {
	case result = <-channel1:
		fmt.Println(result)
	case result = <-channel2:
		fmt.Println(result)
	}
}
