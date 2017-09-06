package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type MeetUp struct {
	team string
	text string
	href string
	when string
	time string
}

func main() {

	list := make(map[string]string)

	doc, _ := goquery.NewDocument("https://www.meetup.com/ko-KR/find/events/tech/?allMeetups=false&radius=26&userFreeform=Seoul%2C+%ED%95%9C%EA%B5%AD+(%EB%8C%80%ED%95%9C%EB%AF%BC%EA%B5%AD)&mcName=%EC%84%9C%EC%9A%B8%2C+KR&lat=37.511093&lon=126.974304.kr/")
	doc.Find(".searchResults").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Each(func(i int, s *goquery.Selection) {

			//li := s.Find(".date-indicator").Text()
			//fmt.Println(li)
			li := s.Find(".event-listing-container")

			li.Find("div").Each(func(i int, s *goquery.Selection) {

				link := s.Find("a").AttrOr("href", "")
				team := s.Find("span").AddClass("itemprop", "name").Text()

				if link != "" && team != "" {
					list[team] = link
				}
			})

		})
	})

	for key, val := range list {
		fmt.Println("key : ", key, " val : ", val)
	}

}
