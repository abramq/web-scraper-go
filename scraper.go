package main

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	cOpis := ""
	cOpisNew1 := ""
	cOpisNew2 := ""
	
	// On every a element which has href attribute call callback
	c.OnHTML("section#count div.container div.row div.col-md-3", func(e *colly.HTMLElement) {	
		cOpis = e.Text

		reg, _ := regexp.Compile("[^a-zA-Z]+ ")	
		cOpisNew1 = reg.ReplaceAllString(cOpis, "")
		cOpisNew1 = strings.Trim(cOpisNew1, " ")
		fmt.Println("txt->",cOpisNew1,"<-")

		reg, _ = regexp.Compile("[^0-9]+")	
		cOpisNew2 = reg.ReplaceAllString(cOpis, "")
		cOpisNew2 = strings.Trim(cOpisNew2, " ")
		fmt.Println("num->",cOpisNew2,"<-")
	})	
	
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	// Start scraping on https://www.mosir.zgora.pl/
	c.Visit("https://www.mosir.zgora.pl/")
}