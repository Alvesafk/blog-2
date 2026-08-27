-- name: CreatePost :one
INSERT INTO posts(
	slug_title,
	title,
	body,
	created_at,
	updated_at,
	tags
)
VALUES(
	$1,
	$2,
	$3,
	NOW(),
	NOW(),
	$4
)
RETURNING *;

-- name: DeletePostByID :exec
DELETE FROM posts
WHERE id = $1;

-- name: DeletePostBySlugTitle :exec
DELETE FROM posts
WHERE slug_title = $1;

-- name: GetAllPosts :many
SELECT *
FROM posts
ORDER BY created_at DESC;

-- name: GetPostByID :one
SELECT *
FROM posts
WHERE id = $1;

-- name: GetPostBySlugTitle :one
SELECT *
FROM posts
WHERE slug_title = $1;

-- name: UpdatePost :exec
UPDATE posts
SET
	slug_title = $2,
	title = $3,
	body = $4,
	update_at = NOW(),
	tags = $5
WHERE id = $1;

-- Posts schema so i don't get lost while writing the queries.
--posts(
--	slug_title TEXT PRIMARY KEY,
--	title TEXT NOT NULL,
--	created_at TIMESTAMP NOT NULL,
--	updated_at TIMESTAMP NOT NULL,
--	tags TEXT []NOT NULL
--)
