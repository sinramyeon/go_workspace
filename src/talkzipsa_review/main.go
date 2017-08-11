package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("https://play.google.com/store/apps/details?id=com.interpark.shop&hl=ko")

	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".single-review").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("span").After("div").Text()
		fmt.Printf(band)
	})

	doc.Find(".review-text").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title

		fmt.Println(s.Text())
	})
}

func main() {
	ExampleScrape()
}
