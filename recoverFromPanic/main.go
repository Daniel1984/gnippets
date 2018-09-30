package main

import "fmt"

func main() {
	actionThatCausesPanic()
	fmt.Println("What does not kill us makes us stronge!")
}

func actionThatCausesPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic:", r)
		}
	}()
	performOperationThatPanics()
}

func performOperationThatPanics() {
	panic("Lets panic!")
}
