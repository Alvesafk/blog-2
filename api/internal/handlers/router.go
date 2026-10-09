package handlers

import (
	"github.com/Alvesafk/blog-2/api/internal/database"
	"github.com/gin-gonic/gin"
)

func NewRouter(conn *database.Connection) *gin.Engine {
	h := &Handler{conn: conn}

	r := gin.Default()

	r.GET("/ping", h.pong)

	return r
}
