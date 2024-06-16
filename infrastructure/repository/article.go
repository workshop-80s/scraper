package repository

import (
	"fmt"
	"scraper/entity"
	"scraper/infrastructure/storage/mysql"
	"scraper/infrastructure/storage/mysql/model"
)

type (
	article struct {
		artisan *mysql.MySql
	}
)

func NewArticle(artisan *mysql.MySql) article {
	return article{
		artisan: artisan,
	}
}

func (r *article) Save(e entity.Article) error {
	m := model.NewArticle(
		e.ID,
		e.SourceID,
		e.Title,
		e.Caption,
		e.Content,
		e.OriginalUrl,
		e.PostTime,
	)
	fmt.Printf("%+v", m)
	r.artisan.Conn.Create(&m)

	return nil
}
