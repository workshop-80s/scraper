package article

import (
	_ "bytes"
	_ "encoding/json"
	"scraper/infrastructure/repository"
	cafef "scraper/infrastructure/source/cafef"

	"scraper/infrastructure/storage/mysql"
)

func Crawl() {
	database := mysql.NewMySql()
	database.Connect()
	defer database.Disconnect()

	repo := repository.NewArticle(database)

	// c := cafef.NewCafeF()
	// c.Crawl()

	// e := entity.NewArticle(
	// 	0,
	// 	10,
	// 	"title",
	// 	"caption",
	// 	"content",
	// 	"url",
	// 	time.Now(),
	// )

	const url = "https://cafef.vn/goc-nhin-chuyen-gia-nhip-dieu-chinh-la-dieu-kien-can-de-buoc-vao-song-moi-mot-so-nhom-co-phieu-co-the-bat-tang-som-hon-thi-truong-188240616123341496.chn"
	e := cafef.CrawlDetail(url)

	
	repo.Save(e)
}
