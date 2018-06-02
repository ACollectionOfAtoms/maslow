package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ACollectionOfAtoms/maslow"
	"github.com/gocolly/colly"
)

func ppT(l []*maslow.Therapist) {
	fmt.Println("[")
	lastIndex := len(l) - 1
	for i, t := range l {
		b, err := json.MarshalIndent(t, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b[:]))
		if i != lastIndex {
			fmt.Println(",")
		}
	}
	fmt.Println("]")
}

func main() {
	c := colly.NewCollector()
	therapists := []*maslow.Therapist{}

	c.OnHTML(".result-row", func(e *colly.HTMLElement) {
		name := strings.Split(e.Attr("data-prof-name"), " ")
		firstName := name[0]
		lastName := strings.Join(name[1:], " ")
		phoneNumber := e.Attr("data-phone")
		imageURL := e.Attr("data-profile-image")

		t := maslow.Therapist{
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: phoneNumber,
			ImageURL:    imageURL,
		}
		therapists = append(therapists, &t)
	})
	c.OnScraped(func(r *colly.Response) {
		ppT(therapists)
	})
	c.Visit("https://www.psychologytoday.com/us/therapists?search=austin")
}
