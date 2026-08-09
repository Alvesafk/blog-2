package db

import (
	"github.com/Alvesafk/blog-2/back/internal/models"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, title, preview, content string, tags []string) (int, error) {
	post := models.Post{
		Title:   title,
		Preview: preview,
		Content: content,
		Tags:    tags,
	}

	result := db.Create(&post)
	if result.Error != nil {
		return -1, result.Error
	}

	return int(result.RowsAffected), nil
}

func GetPostByID(db *gorm.DB, id int) (*models.Post, error) {
	var post models.Post

	if err := db.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPostByTitle(db *gorm.DB, title string) (*models.Post, error) {
	var post models.Post

	if err := db.First(&post, models.Post{Title: title}).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func GetAllPosts(db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func GetMostRecentPost(db *gorm.DB) (*models.Post, error) {
	var post models.Post

	if err := db.Last(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func UpdatePost(db *gorm.DB, postID int, updatedPost models.Post) (*models.Post, error) {
	post, err := GetPostByID(db, postID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(post).Updates(updatedPost).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func DeletePost(db *gorm.DB, post models.Post) error {
	if err := db.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

func DeletePostByID(db *gorm.DB, postID int) error {
	if err := db.Delete(&models.Post{}, postID).Error; err != nil {
		return err
	}

	return nil
}

func CreateComment(db *gorm.DB, content, author string, postID int) (int, error) {
	// Checking if post exists.
	var post models.Post
	if err := db.First(&post, postID).Error; err != nil {
		return -1, err
	}

	comment := models.Comment{
		Content: content,
		Author:  author,
		PostID:  uint(postID),
	}

	result := db.Create(&comment)
	if result.Error != nil {
		return -1, result.Error
	}

	return int(result.RowsAffected), nil
}

func GetCommentsByPostID(db *gorm.DB, postID int) ([]models.Comment, error) {
	// Checking if post exists.
	var post models.Post
	if err := db.First(&post, postID).Error; err != nil {
		return nil, err
	}

	var comments []models.Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentByID(db *gorm.DB, id int) (*models.Comment, error) {
	var comment models.Comment

	if err := db.First(&comment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func UpdateComment(db *gorm.DB, commentID int, updatedComment models.Comment) (*models.Comment, error) {
	comment, err := GetCommentByID(db, commentID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(comment).Updates(updatedComment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func DeleteComment(db *gorm.DB, comment models.Comment) error {
	if err := db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}

func DeleteCommentByID(db *gorm.DB, commentID int) error {
	if err := db.Delete(&models.Comment{}, commentID).Error; err != nil {
		return err
	}

	return nil
}
