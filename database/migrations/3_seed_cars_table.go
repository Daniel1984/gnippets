package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Seeding cars table")
		_, err := db.Exec(`INSERT INTO cars (make, model) VALUES ('Audi', 'RS6')`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("Truncating cars table")
		_, err := db.Exec(`TRUNCATE cars`)
		return err
	})
}
