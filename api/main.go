package main

import "github.com/Alvesafk/blog-2/api/internal/server"

func main() {
	srv := server.New()

	srv.Run()
}
