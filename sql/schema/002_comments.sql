-- +goose Up
CREATE TABLE comments(
	id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	author TEXT NOT NULL,
	body TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	post_id INTEGER NOT NULL CONSTRAINT fk_post_id REFERENCES posts(id)
);
-- +goose Down
DROP TABLE comments;
