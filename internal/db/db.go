package db

import (
	"fmt"
	"os"

	"github.com/Alvesafk/blog-2/back/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn, err := getDSNString()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.Current{}); err != nil {
		return nil, err
	}

	return db, nil
}

func getDSNString() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err 
	}

	d := map[string]string{
		"DB_HOST":     "",
		"DB_USER":     "",
		"DB_PASSWORD": "",
		"DB_NAME":     "",
		"DB_PORT":     "",
		"DB_SSL_MODE": "",
	}

	for k := range d {
		d[k] = os.Getenv(k)
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		d["DB_HOST"], d["DB_USER"], d["DB_PASSWORD"], d["DB_NAME"], d["DB_PORT"], d["DB_SSL_MODE"]), nil
}
