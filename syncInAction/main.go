package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func(i int) {
			fmt.Printf("Waiting for routine %v to end \n", i)
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
}
