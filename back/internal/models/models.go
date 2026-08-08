package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string         `json:"title"`
	Preview   string         `json:"preview"`
	Content   string         `json:"content"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`
}

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	Author  string `json:"author"`
	PostID  uint
	Post    Post
}

type Current struct {
	gorm.Model
	Content string `json:"content"`
}
