// 웹에서 기사 가져오기

// 참고
//// rss
/*
http://readme.skplanet.com
http://www.popit.kr
https://github.com/mmcdole/gofeed
*/

//// 웹 크롤링
/*
https://github.com/PuerkitoBio/goquery
*/

////

// 블로그에서 가져오기

// 트위터에서 가져오기
/*
https://twitter.com/devsfarm
https://twitter.com/WEIRDxMEETUP
https://twitter.com/golangweekly
https://twitter.com/newsycombinator
*/

// 뉴스 사이트에서 가져오기

// 개발자 커뮤니티에서 가져오기

// 깃허브에서 가져오기

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type slackReturn struct {
	Title string
	URL   string
}

func getQuote(sym string) string {
	sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if len(rows) >= 1 && len(rows[0]) == 5 {
		return fmt.Sprintf("%s (%s) is trading at $%s", rows[0][0], rows[0][1], rows[0][2])
	}
	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
}

func OkkyScrape() []string {
	doc, err := goquery.NewDocument("https://okky.kr/")
	if err != nil {
		log.Fatal(err)
	}

	var okkylist []string

	// .Find(".class tag .class")
	// .Find("a").Text() a 태그 안에 든 거

	doc.Find(".article-middle-block").Each(func(i int, s *goquery.Selection) {

		title := s.Find("h5").Text()
		url := "https://okky.kr" + s.Find("h5 a").AttrOr("href", "없음")
		okkylist = append(okkylist, title, url)

	})

	fmt.Println(okkylist)

	return okkylist

}
