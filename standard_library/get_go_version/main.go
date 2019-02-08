package main

import (
	"log"
	"runtime"
)

const info = `The GO version is: %s`

func main() {
	log.Printf(info, runtime.Version())
}
