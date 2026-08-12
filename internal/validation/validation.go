package validation

import (
	"fmt"
	"strings"

	"github.com/Alvesafk/blog-2/back/internal/models"

	"github.com/microcosm-cc/bluemonday"
)

var (
	richPolicy   = bluemonday.UGCPolicy()
	strictPolicy = bluemonday.StrictPolicy()
)

const (
	maxTitleLen   = 200
	maxPreviewLen = 300
	maxContentLen = 60_000
	maxTagLen     = 30
	maxTagsCount  = 5

	maxAuthorLen  = 100
	maxCommentLen = 2_000
)

func SanitizePost(post models.Post) (*models.Post, error) {
	title := strings.TrimSpace(post.Title)
	preview := strings.TrimSpace(post.Preview)
	content := strings.TrimSpace(post.Content)

	if title == "" {
		return nil, fmt.Errorf("title is required")
	}

	if len(title) > maxTitleLen {
		return nil, fmt.Errorf("title exceeds the max amount of characters")
	}

	if len(preview) > maxPreviewLen {
		return nil, fmt.Errorf("preview exceeds the max amount of characters")
	}

	if content == "" {
		return nil, fmt.Errorf("content is required")
	}

	if len(content) > maxContentLen {
		return nil, fmt.Errorf("content exceeds the max amount of characters")
	}

	title = strictPolicy.Sanitize(post.Title)
	preview = strictPolicy.Sanitize(post.Preview)

	content = richPolicy.Sanitize(post.Content)

	tags, err := sanitizeTags(post.Tags)
	if err != nil {
		return nil, err
	}

	return &models.Post{
		Title:   title,
		Preview: preview,
		Content: content,
		Tags:    tags,
	}, nil
}

func sanitizeTags(tags []string) ([]string, error) {
	if len(tags) > maxTagsCount {
		return nil, fmt.Errorf("too many tags")
	}

	seen := make(map[string]bool)
	clean := make([]string, 0, len(tags))

	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t == "" {
			continue
		}

		if len(t) > maxTagLen {
			return nil, fmt.Errorf("tag %q exceeds max amount of characters", t)
		}

		t = strictPolicy.Sanitize(t)
		if !seen[t] {
			seen[t] = true
			clean = append(clean, t)
		}
	}

	return clean, nil
}

func SanitizeComment(comment models.Comment) (*models.Comment, error) {
	content := strings.TrimSpace(comment.Content)
	author := strings.TrimSpace(comment.Author)

	if content == "" {
		return nil, fmt.Errorf("content is required")
	}

	if len(content) > maxCommentLen {
		return nil, fmt.Errorf("content exceeds max amount of characters")
	}

	if author == "" {
		return nil, fmt.Errorf("author is required")
	}

	if len(author) > maxAuthorLen {
		return nil, fmt.Errorf("author exceeds max amount of characters")
	}

	content = strictPolicy.Sanitize(content)
	author = strictPolicy.Sanitize(author)

	return &models.Comment{
		Content: content,
		Author:  author,
	}, nil
}
