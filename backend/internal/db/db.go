package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	conn *sql.DB
}

func New(connString string) (*DB, error) {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	db := &DB{conn}
	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("error on migrations: %s", err)
	}

	return db, nil
}

func (db *DB) migrate() error {
	sourceDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("error on loading migration files: %s", err)
	}

	dbDriver, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error on creating db driver: %s", err)
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return fmt.Errorf("error on initializing migrate: %s", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error on making migrations: %s", err)
	}

	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
