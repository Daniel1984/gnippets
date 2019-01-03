package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	writeToFile("Hello world")
	readFromoFile()
}

func writeToFile(msg string) {
	bytes := []byte(msg)
	ioutil.WriteFile("./test.txt", bytes, 0644)
	fmt.Println("Done writing to file")
}

func readFromoFile() {
	data, _ := ioutil.ReadFile("./test.txt")
	fmt.Printf("File content: %v", string(data))
}
