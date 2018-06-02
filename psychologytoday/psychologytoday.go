package main

import (
	"fmt"

	"github.com/ACollectionOfAtoms/maslow"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(".result-row", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("data-prof-name"))
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

	c.Visit("https://www.psychologytoday.com/us/therapists?search=austin")
}
