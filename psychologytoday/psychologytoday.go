package main

import (
	"fmt"

	"github.com/ACollectionOfAtoms/maslow"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		t := maslow.Therapist{
			FirstName: "Adam",
			LastName:  "Hernandez",
		}
		fmt.Println(t)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}
