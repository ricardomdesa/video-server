package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:bR632F9EYioo@localhost:5432/course?sslmode=enable")
	if err!= nil {
        panic(err)
    }
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"postgres", driver)
	if err!= nil {
        panic(err)
    }
	m.Steps(2)
	// m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
