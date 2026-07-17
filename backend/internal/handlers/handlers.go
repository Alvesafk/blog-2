package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sort"
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

func (r Response) Write(w http.ResponseWriter, h int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(r)
}

func (s *Connection) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.ListPosts()
	if err != nil {
		http.Error(w, "could not get posts", http.StatusInternalServerError)
		return
	}

	if len(posts) < 1 {
		Response{
			Message: "There is no post",
			Status:  "Failed",
		}.Write(w, http.StatusNotFound)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ID > posts[j].ID
	})

	Response{
		Message: "Success",
		Status:  "ok",
		Content: posts,
	}.Write(w, http.StatusOK)
}

func (s *Connection) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		Response{
			Message: "Invalid id",
			Status:  "Failed",
		}.Write(w, http.StatusBadRequest)
		return
	}

	post, err := s.db.GetPostByID(id)
	if err != nil {
		Response{
			Message: "Post does not exist",
			Status:  "Failed",
		}.Write(w, http.StatusNotFound)
		return
	}

	post.Content = mdToHtml(post.Content)

	Response{
		Message: "Success",
		Status:  "ok",
		Content: post,
	}.Write(w, http.StatusOK)
}

func (s *Connection) GetLatestPost(w http.ResponseWriter, r *http.Request) {
	post, err := s.db.GetLatestPost()
	if err != nil {
		if err.Error() == "latest post was not found" {
			Response{
				Message: "There is no post",
				Status:  "Failed",
			}.Write(w, http.StatusNotFound)
			return

		}

		Response{
			Message: "Could not get latest post",
			Status:  "Failed",
		}.Write(w, http.StatusInternalServerError)
		return
	}

	Response{
		Message: "Success",
		Status:  "ok",
		Content: post,
	}.Write(w, http.StatusOK)
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
		http.Error(w, "could not get comments", http.StatusNotFound)
		return
	}

	Response{
		Message: "Success",
		Status:  "ok",
		Content: comments,
	}.Write(w, http.StatusOK)
}

func (s *Connection) HealthCheck(w http.ResponseWriter, r *http.Request) {
	Response{
		Message: "Success",
		Status:  "healthy",
	}.Write(w, http.StatusOK)
}

func mdToHtml(md string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
