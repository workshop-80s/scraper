package model

import "time"

type Article struct {
	ID       int    `gorm:"primary_key;column:id"`
	SourceID int    `gorm:"column:source_id"`
	Title    string `gorm:"column:title"`
	// Caption     string    `gorm:"column:caption"`
	Content     string `gorm:"column:content"`
	OriginalUrl string `gorm:"column:url"`
	// PostTime    time.Time `gorm:"column:post_time"`
}

func NewArticle(
	id int,
	sourceId int,
	title,
	caption,
	content,
	originalUrl string,
	postTime time.Time,
) Article {
	return Article{
		ID:       id,
		SourceID: sourceId,
		Title:    title,
		// Caption:     caption,
		Content:     content,
		OriginalUrl: originalUrl,
		// PostTime:    postTime,
	}
}
