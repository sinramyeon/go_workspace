package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("https://okky.kr/")
	if err != nil {
		log.Fatal(err)
	}

	// .Find(".class tag .class")
	// .Find("a").Text() a 태그 안에 든 거

	doc.Find(".article-middle-block").Each(func(i int, s *goquery.Selection) {

		title := s.Find("h5").Text()

		url := "https://okky.kr" + s.Find("h5 a").AttrOr("href", "없음")

		println(title)
		println(url)

	})
}

func main() {
	ExampleScrape()
}
