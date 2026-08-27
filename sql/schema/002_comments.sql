-- +goose Up
CREATE TABLE comments(
	id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	author TEXT NOT NULL,
	body TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	post_title TEXT NOT NULL CONSTRAINT fk_post_title REFERENCES posts(slug_title)
);
-- +goose Down
DROP TABLE comments;
