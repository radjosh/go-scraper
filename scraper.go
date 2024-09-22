package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.scrapingcourse.com/ecommerce/"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var productNames []string
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		productNames.append(productNames, s.Text())
	})

	for _, name := range productNames {
		fmt.Println(name)
	}
}