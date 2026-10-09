package main

import (
	"context"
	"log"

	"github.com/Alvesafk/blog-2/api/internal/server"
)

func main() {
	ctx := context.Background()

	srv, err := server.New(ctx)
	if err != nil {
		log.Println("Error: could not create a server config:", err)
		return
	}

	srv.Run()
}
