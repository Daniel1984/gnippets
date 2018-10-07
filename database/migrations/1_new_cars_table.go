package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Creating cars table")
		_, err := db.Exec(`CREATE TABLE cars(id SERIAL PRIMARY KEY, make varchar(50) NOT NULL)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping cars table")
		_, err := db.Exec(`DROP TABLE cars`)
		return err
	})
}
