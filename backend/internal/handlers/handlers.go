package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alvesafk/blog-2/internal/db"
)

type Connection struct {
	db *db.DB
}

func NewConnection(db *db.DB) *Connection {
	return &Connection{db: db}
}


type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Content any    `json:"content"`
}

func (r Response) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}

func (s *Connection) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Response{
		Message: "Success",
		Status:  "ok",
		Content: posts,
	}.Write(w)
}

func (s *Connection) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	post, err := s.db.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	Response{
		Message: "Success",
		Status:  "ok",
		Content: post,
	}.Write(w)
}

func (s *Connection) GetComments(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	comments, err := s.db.ListCommentsByPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	Response{
		Message: "Success",
		Status: "ok",
		Content: comments,
	}.Write(w)
}

func (s *Connection) HealthCheck(w http.ResponseWriter, r *http.Request) {
	Response{
		Message: "Success",
		Status:  "healthy",
		Content: "",
	}.Write(w)
}
