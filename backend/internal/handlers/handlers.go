package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alvesafk/blog-2/internal/db"
	"github.com/yuin/goldmark"
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
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(r)
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

func (s *Connection) GetLatestPost(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	latestPost := posts[len(posts) - 1]
	latestPost.Content = mdToHtml(latestPost.Content)

	Response{
		Message: "Success",
		Status: "ok",
		Content: latestPost,
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

func mdToHtml(md string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
