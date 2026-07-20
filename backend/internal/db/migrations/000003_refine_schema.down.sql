ALTER TABLE currently
	DROP CONSTRAINT IF EXISTS currently_content_not_empty;

ALTER TABLE comments
	DROP CONSTRAINT IF EXISTS comments_content_not_empty;

ALTER TABLE comments
	DROP CONSTRAINT IF EXISTS comments_author_not_empty;

DROP INDEX IF EXISTS idx_comments_post_id;

ALTER TABLE posts
	DROP CONSTRAINT IF EXISTS posts_title_not_empty;

ALTER TABLE posts
	DROP CONSTRAINT IF EXISTS posts_slug_title_unique;

ALTER TABLE posts
	DROP COLUMN IF EXISTS updated_at;
