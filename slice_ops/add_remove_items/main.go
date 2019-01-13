package main

import (
	"fmt"
)

func main() {
	addElementsToSpecificPlace([]int{4, 5, 6}, 3)
	expandSlice()
}

func addElementsToSpecificPlace(b []int, i int) {
	a := []int{1, 2, 3, 7, 8, 9}
	a = append(a[:i], append(b, a[i:]...)...)
	fmt.Println(a)
}

func expandSlice() {
	months := []string{"January", "February", "December"}
	months = append(months[:2], append(make([]string, 9), months[2:]...)...)
	fmt.Println(months)
}
