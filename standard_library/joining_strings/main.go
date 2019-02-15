package main

import (
	"fmt"
	"strings"
)

func sampleOne() {
	selectBase := "SELECT * FROM user WHERE %s "
	refStringSlice := []string{
		" FIRST_NAME = 'Jack' ",
		" INSURANCE_NO = 333444555 ",
		" EFFECTIVE_FROM = SYSDATE ",
	}

	sentence := strings.Join(refStringSlice, "AND")
	fmt.Printf(selectBase+"\n", sentence)
}

func main() {
	sampleOne()
}
