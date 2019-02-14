package main

import (
	"fmt"
	"strings"
)

func breakOnSpace() {
	str := "Quick Brown Fox"

	words := strings.Fields(str)
	for i, word := range words {
		fmt.Printf("Word %d is: %s\n", i, word)
	}
}

func main() {
	breakOnSpace()
}
