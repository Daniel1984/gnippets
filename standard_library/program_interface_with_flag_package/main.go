package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom type need to implement
// flag.Value interface to be able to
// use it in flag.Var function.
type Values []string

func (s *Values) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *Values) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	// Extracting flag values with methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")
	arr := &Values{}
	flag.Var(arr, "array", "Input array to iterate through.")
	/* Execute the flag.Parse function, to
	* read the flags to defined variables.
	* Without this call the flag
	* variables remain empty.
	 */
	flag.Parse()
	// Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Array of args %v\n", arr)
		retryCount++
	}
}
