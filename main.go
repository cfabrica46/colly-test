package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Link struct {
	Text      string `json:"text"`
	Link      string `json:"link"`
	ClassName string `json:"className"`
}

func newLink(text, link, class string) Link {
	return Link{Text: text, Link: link, ClassName: class}
}

func main() {
	var links []Link
	c := colly.NewCollector()

	c.OnHTML("a.gb1", func(element *colly.HTMLElement) {
		link := newLink(element.Text, element.Attr("href"), element.Attr("class"))
		links = append(links, link)
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Printf("Visiting %s...\n\n", request.URL.String())
	})

	// c.Visit("https://hectormainar.com/h1.php")
	err := c.Visit("https://www.google.com.pe/")
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(links, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
