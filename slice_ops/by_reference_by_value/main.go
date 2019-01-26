package main

import "fmt"

func main() {
	array := [3]string{"one", "two", "three"}
	arrayByVal := array
	arrayByVal[1] = "by val"

	fmt.Println("expecting two:", array[1])

	arrayByRef := &array
	arrayByRef[1] = "by ref"
	fmt.Println("expecting by ref:", array[1])

	byRef(&array)
	fmt.Println("expecting by ref:", array[0])
}

func byRef(arr *[3]string) {
	arr[0] = "by ref"
}
