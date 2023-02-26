package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn string
}

func NewDatabase(d Database) (*sql.DB, error) {
	// Connect to database
	db, err := sql.Open("postgres", d.Conn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
