package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fileName := "data.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file! err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// find and print all links, write to .txt

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, str := range links {
			fmt.Printf("\n%s\n", str)
		}

	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")

	table()
}
