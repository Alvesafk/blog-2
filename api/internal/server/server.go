package server

import (
	"github.com/Alvesafk/blog-2/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	return &Server{
		router: handlers.NewRouter(),
	}
}

func (srv *Server) Run() {
	srv.router.Run()
}
