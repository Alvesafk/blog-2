package handlers

import (
	"github.com/Alvesafk/blog-2/api/internal/database"
)

type Handler struct {
	conn *database.Connection
}
