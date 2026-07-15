package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID       int            `json:"id"`
	Title    string         `json:"title"`
	Content  string         `json:"content"`
	PostedAt time.Time      `json:"postedAt"`
	Tags     pq.StringArray `json:"tags"`
}

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"postId"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	CommentedAt time.Time `json:"commentedAt"`
}

func (db *DB) InsertPost(title, content string, tags []string) (int, error) {
	var id int
	err := db.conn.QueryRow(
		`INSERT INTO posts (title, content, tags) VALUES ($1, $2, $3) RETURNING id`,
		title, content, pq.Array(tags),
	).Scan(&id)

	return id, err
}

func (db *DB) ListPosts() ([]Post, error) {
	rows, err := db.conn.Query(`SELECT id, title, content, posted_at, tags FROM posts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.PostedAt, &p.Tags); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (db *DB) GetPostByID(id int) (*Post, error) {
	var p Post
	err := db.conn.QueryRow(
		`SELECT id, title, content, posted_at, tags FROM posts WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.Title, &p.Content, &p.PostedAt, &p.Tags)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post %d not found", id)
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (db *DB) UpdatePost(id int, new_post Post) error {
	_, err := db.conn.Exec(
		`UPDATE posts SET title = $1, content = $2, tags = $3 WHERE id = $4`,
		new_post.Title, new_post.Content, pq.Array(new_post.Tags), id,
	)

	return err
}

func (db *DB) DeletePost(id int) error {
	_, err := db.conn.Exec(`DELETE FROM posts WHERE id = $1`, id)
	return err
}

func (db *DB) InsertComment(post_id int, content, author string) (int, error) {
	var id int
	err := db.conn.QueryRow(
		`INSERT INTO comments (post_id, content, author) VALUES ($1, $2, $3) RETURNING id`,
		post_id, content, author,
	).Scan(&id)

	return id, err
}

func (db *DB) ListCommentsByPost(post_id int) ([]Comment, error) {
	rows, err := db.conn.Query(
		`SELECT id, post_id, content, author, commented_at FROM comments WHERE post_id = $1`,
		post_id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.Content, &c.Author, &c.CommentedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, rows.Err()
}

func (db *DB) GetCommentByID(id int) (*Comment, error) {
	var c Comment
	err := db.conn.QueryRow(
		`SELECT id, post_id, content, author, commented_at FROM comments WHERE id = $1`,
		id,
	).Scan(&c.ID, &c.PostID, &c.Content, &c.Author, &c.CommentedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("comment %d not found", id)
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (db *DB) UpdateComment(id int, new_comment Comment) error {
	_, err := db.conn.Exec(
		`UPDATE comments SET content = $1, author = $2 WHERE id = $3`,
		new_comment.Content, new_comment.Author, id,
	)

	return err
}

func (db *DB) DeleteComment(id int) error {
	_, err := db.conn.Exec(`DELETE FROM comments WHERE id = $1`, id)
	return err
}
