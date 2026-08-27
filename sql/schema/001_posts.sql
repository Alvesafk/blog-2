-- +goose Up
CREATE TABLE posts(
	slug_title TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	tags TEXT []NOT NULL
);
-- +goose Down
DROP TABLE posts;
