// db/migrate_up.go
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cmd := os.Args[1]
	db, err := sql.Open("sqlite3", "./videos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./migrations/migrations_sqlite")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	if cmd == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
        if err := m.Down(); err != nil {
            log.Fatal(err)
        }
    }	
}