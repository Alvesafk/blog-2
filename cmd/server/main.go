package main

import (
	"fmt"

	"github.com/Alvesafk/blog-2/back/internal/db"
)

func main() {
	d, err := db.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected with the db!")
	fmt.Println(d.Name())
}
