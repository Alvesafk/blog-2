-- posts: garantir unicidade de slug e rastrear atualizações
ALTER TABLE posts
	ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT now();

ALTER TABLE posts
	ADD CONSTRAINT posts_slug_title_unique UNIQUE (slug_title);

ALTER TABLE posts
	ADD CONSTRAINT posts_title_not_empty CHECK (length(trim(title)) > 0);

-- comments: indexar FK e impedir conteúdo vazio
CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);

ALTER TABLE comments
	ADD CONSTRAINT comments_author_not_empty CHECK (length(trim(author)) > 0);

ALTER TABLE comments
	ADD CONSTRAINT comments_content_not_empty CHECK (length(trim(content)) > 0);

-- currently: impedir conteúdo vazio
ALTER TABLE currently
	ADD CONSTRAINT currently_content_not_empty CHECK (length(trim(content)) > 0);
