package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func New(connString string) (*DB, error) {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	db := &DB{conn}
	return db, db.migrate()
}

func (db *DB) migrate() error {
	_, err := db.conn.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			posted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			tags TEXT[]
		);

		CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
			content TEXT NOT NULL,
			author TEXT NOT NULL,
			commented_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);
	`)

	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}
