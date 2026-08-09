package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alvesafk/blog-2/back/internal/db"
	"github.com/Alvesafk/blog-2/back/internal/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()

	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	conn := handlers.NewConnection(d)
	r := defRouter(*conn)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err = run(ctx, server); err != nil {
		log.Fatal(err)
	}
}

func defRouter(conn handlers.Connection) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/hello", conn.GetHelloWorld)

	return r
}

func run(ctx context.Context, s *http.Server) error {
	serverError := make(chan error, 1)

	go func() {
		log.Printf("Server is running on http//localhost%s", s.Addr)
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		return err
	case sig := <-stop:
		fmt.Println()
		log.Printf("Received shutdown signal: %v", sig)
	}

	log.Println("Server is shutting down...")

	ctxShutdown, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctxShutdown); err != nil {
		log.Printf("Server shutdown error: %s", err)
		log.Printf("Using os.Exit(1)...")
		os.Exit(1)
	}

	log.Println("Server exited properly.")

	return nil
}
