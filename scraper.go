package main

import (
	"fmt"
	"os"
	"github.com/anaskhan96/soup"
)

func main() {
	url := "https://www.scrapingcourse.com/ecommerce/"

	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	productNames := doc.Find("div").FindAll("h2")
	for _, name := range productNames {
		fmt.Println(name.Text())
	}
}