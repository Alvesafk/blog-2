package server

import (
	"context"

	"github.com/Alvesafk/blog-2/api/internal/database"
	"github.com/Alvesafk/blog-2/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	conn   *database.Connection
}

func New(ctx context.Context) (*Server, error) {
	conn, err := database.Open("blog.db")
	if err != nil {
		return nil, err
	}

	return &Server{
		router: handlers.NewRouter(conn),
		conn:   conn,
	}, nil
}

func (srv *Server) Run() {
	srv.router.Run()
}
