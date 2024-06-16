package entity

import "time"

type Article struct {
	ID          int
	SourceID    int
	Title       string
	Caption     string
	Content     string
	OriginalUrl string
	PostTime    time.Time
}

func NewArticle(
	id,
	sourceId int,
	title,
	caption,
	content,
	originalUrl string,
	postTime time.Time,
) Article {
	return Article{
		ID:          id,
		SourceID:    sourceId,
		Title:       title,
		Caption:     caption,
		Content:     content,
		OriginalUrl: originalUrl,
		PostTime:    postTime,
	}
}
