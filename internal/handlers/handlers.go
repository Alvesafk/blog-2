package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection(db *gorm.DB) *Connection {
	return &Connection{db}
}

type Response struct {
	Message string `json:"message"`
	Content any    `json:"content"`
}

func (r Response) WriteJSON(w http.ResponseWriter, h int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	err := enc.Encode(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c Connection) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	Response{
		Message: "This is a hello world.",
		Content: "Hello, World!",
	}.WriteJSON(w, http.StatusOK)
}
