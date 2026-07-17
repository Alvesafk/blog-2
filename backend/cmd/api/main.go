package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alvesafk/blog-2/internal/db"
	"github.com/Alvesafk/blog-2/internal/handlers"
	mw "github.com/Alvesafk/blog-2/internal/middlewares"
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

	con := handlers.NewConnection(database)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/post", con.GetPosts)
	mux.HandleFunc("GET /api/post/{id}", con.GetPost)
	mux.HandleFunc("GET /api/post/latest", con.GetLatestPost)
	mux.HandleFunc("GET /api/comment/{id}", con.GetComments)
	mux.HandleFunc("GET /api/currently", con.GetCurrently)

	mux.HandleFunc("GET /healthz", con.HealthCheck)

	middlewareMux := mw.Chain(mux,
		mw.RecoverMiddleware,
		mw.LoggingMiddleware,
		mw.CORSMiddleware,
		mw.SecurityHeadersMiddleware,
		mw.RateLimiterMiddleware,
	)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           middlewareMux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		fmt.Println("Listening on port :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error in booting the server: %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
