package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"tools.cyberkrypts.dev/env"
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

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS send_files (
			id SERIAL PRIMARY KEY,
			file_name TEXT NOT NULL,
			file_size INT NOT NULL,
			file_id TEXT NOT NULL,
			web_rtc_session_id TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
}

func GetDb() (*sql.DB, error) {
	return connect()
}

func connect() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}

	DatabaseURL := env.GetEnv().DatabaseURL

	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}

	database = db

	return db, nil
}
