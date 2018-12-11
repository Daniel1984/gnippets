package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"os"
)

var (
	Db   = "gnippertsrest"
	User = "localhost:8181"
)

type ServiceConfig struct {
	Db   string `json:"db"`
	User string `json:"user"`
}

func GetConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{Db, User}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}

func DbConnectionEstablished(db *pg.DB) error {
	var n int
	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		return err
	}

	return nil
}
