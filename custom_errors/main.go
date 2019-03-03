package main

import "fmt"

type myCustomError struct {
	shortMessage string
	description  string
}

func (mce *myCustomError) Error() string {
	errorMsg := fmt.Sprintf("short msg: %v, description: %v", mce.shortMessage, mce.description)
	return errorMsg
}

func funcThatReturnsError() error {
	// some computations that might end up as exception
	return &myCustomError{
		shortMessage: "Error happened",
		description:  "funcThatReturnsError had and exception, please investigate",
	}
}

func main() {
	err := funcThatReturnsError()
	if err != nil {
		fmt.Println(err)
	}
}
