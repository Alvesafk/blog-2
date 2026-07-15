package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Alvesafk/blog-2/internal/db"
	"github.com/Alvesafk/blog-2/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env wasn't loaded: ", err)
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL wasn't defined.")
	}

	database, err := db.New(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	srv := handlers.NewServer(database)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /post", srv.GetPosts)
	mux.HandleFunc("GET /post/{id}", srv.GetPost)

	http.ListenAndServe(":8080", mux)
}
