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

func breakOnSpecialChar() {
	str := "Quick_Brown Fox Jumper_Over_The Lazy Dog"
	words := strings.Split(str, "_")

	for i, word := range words {
		fmt.Printf("Word %d is: %s\n", i, word)
	}
}

func main() {
	breakOnSpace()
	fmt.Println("-------------------------")
	breakOnSpecialChar()
}
