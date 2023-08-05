package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	scrappedUrl := "https://www.imdb.com/chart/toptv"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}

		fmt.Println(r.Method)
	})

	c.OnHTML("h1.ipc-title__text", func(e *colly.HTMLElement) {
		fmt.Println("Category:", e.Text)
	})
	// name of shows
	c.OnHTML("ul li div div div div h3", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visiting:", r.Request.URL, "StatusCode:", r.StatusCode)
	})

	c.Visit(scrappedUrl)
}
