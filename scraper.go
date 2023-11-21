package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".col-md-3 span", func(e *colly.HTMLElement) {		
		fmt.Printf("This found: %q -> %s\n", e.Text )
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	// Start scraping on https://www.mosir.zgora.pl/
	c.Visit("https://www.mosir.zgora.pl/")
}