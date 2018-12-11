package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log_file, err := os.Create("log-file_" + time.RFC3339)

	if err != nil {
		log.Fatal("Can't create log file")
	}

	defer log_file.Close()

	log.SetOutput(log_file)
	log.Fatalln("Log: Something went wrong!")
	log.Println("Extra logging here...")
}
