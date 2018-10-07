package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("Adding model to cars table")
		_, err := db.Exec(`ALTER TABLE cars ADD model varchar(50) NOT NULL`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("Removing model colmn from cars table")
		_, err := db.Exec(`ALTER TABLE cars DROP model`)
		return err
	})
}
