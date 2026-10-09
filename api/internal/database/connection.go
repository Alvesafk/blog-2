package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	openDBError    = errors.New("couldn't open the DB")
	migrationError = errors.New("couldn't migrate the models onto the DB")
)

type Connection struct {
	db *gorm.DB
}

func Open(dbName string) (*Connection, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, openDBError
	}

	if err := migrate(db); err != nil {
		return nil, migrationError
	}

	return &Connection{db: db}, nil
}

// No model yet to migrate so will just return nil, for now!
func migrate(db *gorm.DB) error {
	return nil
}
