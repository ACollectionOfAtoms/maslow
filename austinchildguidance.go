package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Find and grab all therapist photo links
	c.OnHTML("div.collection-item__content", func(e *colly.HTMLElement) {
		imgSrc := e.ChildAttr("img", "src")

		// Print link
		if imgSrc != "" {
			fmt.Printf("Picture found: %s\n", imgSrc)
		}
	})

	// Find and grab all available therapist information
	c.OnHTML("div.collection-item-description", func(e *colly.HTMLElement) {
		therapistInformation := e.ChildText("p")

		// Print link
		if therapistInformation != "" {
			fmt.Printf("Information found: %s\n", therapistInformation)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("https://www.austinchildguidance.org/about-us/board-staff.html")
}
