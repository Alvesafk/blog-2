package database

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"index;not null;unique;size:256"`
	Content string `gorm:"not null"`
	Tags    []Tag  `gorm:"many2many:post_tags;"`
}

type Tag struct {
	gorm.Model
	Name  string `grom:"uniqueIndex;not null"`
	Posts []Post `gorm:"many2many:post_tags;"`
}
