package main

import (
	"fmt"

	"github.com/temorfeouz/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("adding id column...")
		_, err := db.Exec(`ALTER TABLE my_table ADD id serial`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping id column...")
		_, err := db.Exec(`ALTER TABLE my_table DROP id`)
		return err
	})
}
