package main

import "fmt"

func fact(num int) int {
	if num == 0 {
		return 1
	}

	return num * fact(num-1)
}

func main() {
	fmt.Println(fact(5))
}
