package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var database *sql.DB

func Init() {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS short_urls (
			id SERIAL PRIMARY KEY,
			long_url TEXT NOT NULL,
			short_url TEXT NOT NULL UNIQUE
		);
	`)
	if err != nil {
		panic(err)
	}
}

func GetDb() (*sql.DB, error) {
	return connect()
}

func connect() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}

	DB_URL := os.Getenv("DATABASE_URL")
	if DB_URL == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		return nil, err
	}

	database = db

	return db, nil
}
