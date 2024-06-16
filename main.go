package main

import (
	"fmt"
	"log"
	"os"

	"scraper/usecase/article"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "crawl",
				Aliases: []string{"c"},
				Usage:   "crawl a website",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("\ncrawl BEGIN")
					article.Crawl()
					fmt.Println("\ncrawl END")

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
