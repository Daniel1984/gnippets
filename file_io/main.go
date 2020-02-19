package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeToFile("Hello world")
	readFromoFile()
}

func writeToFile(msg string) {
	bytes := []byte(msg)
	ioutil.WriteFile("./test.txt", bytes, 0644)
	fmt.Println("Done writing to file")
	v, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println("err ", err.Error())
	} else {
		fmt.Println(v)
		if err := os.Remove("./test.txt"); err != nil {
			fmt.Println("error delete ", err.Error())
		}
	}
}

func readFromoFile() {
	data, _ := ioutil.ReadFile("./test.txt")
	fmt.Printf("File content: %v", string(data))
}
