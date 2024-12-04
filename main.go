package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	baseUrl := "docs.spring.io"
	errorUrls := make(map[string]string)

	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		urlToVisit := e.Attr("href")
		if !strings.Contains(urlToVisit, baseUrl) {
			return
		}

		errorMessage := e.Request.Visit(urlToVisit)
		if errorMessage != nil && errorMessage.Error() != "URL already visited" {
			// errorUrls[urlToVisit] = errorMessage.Error()
			log.Println("Error Visiting ", urlToVisit)
			log.Println("Error message ", errorMessage.Error())
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ", r.URL)
	})

	builtin := c.Visit("https://docs.spring.io/spring-ldap/reference/spring-ldap-basic-usage.html")
	if builtin != nil {
		log.Fatal(builtin)
	}

	for errorUrl, errorMessage := range errorUrls {
		log.Println(errorUrl, errorMessage)
	}
}
