package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".result-row", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("data-prof-name"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.psychologytoday.com/us/therapists?search=austin")
}
