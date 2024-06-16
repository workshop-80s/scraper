// domain: https://bovagau.vn/explore
package article

import (
	_ "bytes"
	_ "encoding/json"
	"fmt"
	"strings"
	"time"

	"scraper/entity"

	"github.com/gocolly/colly"
)

const LINK = "https://cafef.vn/"
const Domain = "https://cafef.vn"
const sourceID = 1

type cafeF struct {
}

func NewCafeF() cafeF {
	return cafeF{}
}

func (f *cafeF) Crawl() {
	c := f.CrawlCategory()
	fmt.Printf("%v", c)
}

func (f *cafeF) CrawlCategory() []string {
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	result := []string{}

	c.OnHTML(".top_noibat .top_noibat_row1", func(e *colly.HTMLElement) {
		e.ForEach("h2 a", func(i int, e1 *colly.HTMLElement) {
			url := Domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

			result = append(result, url)
		})
	})

	c.OnHTML(".top_noibat .top_noibat_row2", func(e *colly.HTMLElement) {
		e.ForEach("div.big h2 a", func(i int, e1 *colly.HTMLElement) {
			url := Domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")
			result = append(result, url)
		})
	})

	c.OnHTML("#box-nha-dau-tu-default", func(e *colly.HTMLElement) {
		e.ForEach("div.item a", func(i int, e1 *colly.HTMLElement) {
			url := Domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

			result = append(result, url)
		})
	})

	c.OnHTML("#streamItem", func(e *colly.HTMLElement) {
		e.ForEach(".top5_news .listchungkhoannew div.box-category-item div.knswli-right", func(i int, e1 *colly.HTMLElement) {
			e1.ForEach("h3 a", func(i int, e2 *colly.HTMLElement) {
				url := Domain + "/" + strings.TrimLeft(e2.Attr("href"), "/")
				result = append(result, url)
			})
		})
	})

	c.Visit(LINK)

	return result
}

func CrawlDetail(url string) entity.Article {
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	title := ""
	content := ""
	timestamp := ""
	postTime := time.Now()

	c.OnHTML(".totalcontentdetail", func(e *colly.HTMLElement) {

		e.ForEach("h1.title", func(i int, e1 *colly.HTMLElement) {
			title = strings.TrimSpace(e1.Text)
		})

		e.ForEach("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) {
			timestamp = strings.TrimSpace(e1.Text)
		})
	})
	fmt.Println(" ")

	// content
	c.OnHTML("#mainContent", func(e *colly.HTMLElement) {
		e.ForEach(".detail-content > p", func(_ int, e1 *colly.HTMLElement) {
			fmt.Println(" ")
			fmt.Println(e1.Text)

			fmt.Println(" ")

			content += strings.TrimSpace(e1.Text)
		})
	})

	c.Visit(url)

	fmt.Println(timestamp)
	return entity.NewArticle(
		0,
		sourceID,
		title,
		"",
		content,
		url,
		postTime,
	)
}

/*
func detail(url string) {
	c := colly.NewCollector()
	// c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	// result := []Article{}

	c.OnHTML(".totalcontentdetail", func(e *colly.HTMLElement) {
		// title := ""
		// // content := ""
		// timestamp := ""
		// e.ForEach("h1.title", func(i int, e1 *colly.HTMLElement) {
		// 	title = strings.TrimSpace(e1.Text)
		// })

		// e.ForEach("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) {
		// 	timestamp = strings.TrimSpace(e1.Text)
		// })

	})
	fmt.Println(" ")

	// content
	content := ""
	c.OnHTML("#mainContent", func(e *colly.HTMLElement) {
		e.ForEach(".detail-content > p", func(_ int, e1 *colly.HTMLElement) {
			fmt.Println(" ")
			fmt.Println(e1.Text)

			fmt.Println(" ")

			// content += strings.TrimSpace(e1.Text)
		})

	})

	fmt.Println(content)

	c.Visit(url)
}

*/
