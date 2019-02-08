package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println(ex)

	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :" + exPath)

	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("working dir with os.Getwd:", dir)
}
