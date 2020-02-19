package main

import (
	"fmt"
	"os"

	"github.com/gnippets/flags/pkg/apicfg"
)

func main() {
	ac, err := apicfg.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v", ac)
}
