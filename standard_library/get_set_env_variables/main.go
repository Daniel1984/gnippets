package main

import (
	"log"
	"os"
)

func getEnvVariable() {
	key := "DB_CONN"
	connStr, ok := os.LookupEnv(key)

	if !ok {
		log.Printf("The env variable %s is not set.\n", key)
	}

	log.Println(connStr)
}

func setUnsetEnvVariable() {
	key := "DB_CONN"
	defaultValue := "postgres:default"

	if connStr, ok := os.LookupEnv(key); !ok {
		log.Printf("The env variable %s is not set. Setting default value\n", key)
		os.Setenv(key, defaultValue)
	} else {
		log.Printf("The env variable for %s key is %s\n", key, connStr)
		return
	}

	connStr, _ := os.LookupEnv(key)
	log.Printf("Got default variable set %s", connStr)
}

func main() {
	getEnvVariable()
	setUnsetEnvVariable()
}
