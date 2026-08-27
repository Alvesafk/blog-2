-- name: CreateComment :one
INSERT INTO comments (
	author,
	body,
	created_at,
	updated_at,
	post_id
)
VALUES (
	$1,
	$2,
	NOW(),
	NOW(),
	$3
)
RETURNING *;

-- name: DeleteCommentByID :exec
DELETE FROM comments
WHERE id = $1;

-- name: GetAllComments :many
SELECT *
FROM comments;

-- name: GetCommentByID :one
SELECT *
FROM comments
WHERE id = $1;

-- name: GetPostComments :many
SELECT *
FROM comments
WHERE post_id = $1
ORDER BY created_at DESC;

-- name: UpdateComment :exec
UPDATE comments
SET
	author = $2,
	body = $3,
	update_at = NOW()
WHERE id = $1;

-- Comments schema so i don't get lost while writing the queries.
-- comments(
-- 	id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
-- 	author TEXT NOT NULL,
-- 	body TEXT NOT NULL,
-- 	created_at TIMESTAMP NOT NULL,
-- 	updated_at TIMESTAMP NOT NULL,
-- 	post_id INTEGER NOT NULL CONSTRAINT fk_post_id REFERENCES posts(id)
-- );
