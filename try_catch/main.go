package main

import (
	"errors"
	"fmt"
)

func main() {
	if val, err := methodThatThrowns(); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Got response, perform calculation with %v", val)
	}
}

func methodThatThrowns() (string, error) {
	// perform calculations and return value, error
	return "", errors.New("Something happened.")
}
