package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Creating users table...")
		_, err := db.Exec(`
			CREATE TABLE users(
				id SERIAL PRIMARY KEY,
				name varchar(50) NOT NULL,
				role varchar(50) NOT NULL
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping users table...")
		_, err := db.Exec("DROP TABLE users")
		return err
	})
}
