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

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mmcdole/gofeed"
)

type slackReturn struct {
	Title string
	URL   string
}

type twitterConf struct {
	confKey     string
	confSecret  string
	tokenKey    string
	tokenSecret string
}

// 트위터 읽기
func TwitterScrape() map[string]string {

	//1. 트위터에 연결하기
	var env twitterConf
	env = twitterConf(env)
	tweetlist := make(map[string]string)

	// https://apps.twitter.com/app/14097310/keys 를 보시면 됩니다.

	config := oauth1.NewConfig(env.confKey, env.confSecret)
	token := oauth1.NewToken(env.tokenKey, env.confSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// 테스트 하시려면 유저 확인을 해보시면 좋습니다.

	//verifyParams := &twitter.AccountVerifyParams{
	//	SkipStatus:   twitter.Bool(true),
	//	IncludeEmail: twitter.Bool(true),
	//}

	//로그인 유저 얻기
	//_, _, _ := client.Accounts.VerifyCredentials(verifyParams)

	// 특정 계정 크롤링 하기

	tweets := GetUserTweets(2, "golangweekly", client)

	for _, v := range tweets {
		tweetlist[v.Text] = tweetlist[v.Entities.Urls[0].URL]
	}

	tweets = GetUserTweets(2, "WEIRDxMEETUP", client)

	for _, v := range tweets {
		tweetlist[v.Text] = tweetlist[v.Entities.Urls[0].URL]
	}

	tweets = GetUserTweets(2, "devsfarm", client)

	for _, v := range tweets {
		tweetlist[v.Text] = tweetlist[v.Entities.Urls[0].URL]
	}

	// 트윗 내용이 키이고 url이 밸류인 걸로 리턴 했습니다(밸류안에 뭘 넣을지 고민중...)
	return tweetlist

}

// 유저별 트윗을 얻는 모듈
// 얻어오고 싶은 양, 유저 아이디, 클라이언트 접속정보를 넣어 보세요...
func GetUserTweets(many int, id string, client *twitter.Client) []twitter.Tweet {

	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           many,
		ScreenName:      id,
		IncludeRetweets: twitter.Bool(false),
	})

	return tweets

}

// rss 블로그 읽기
func RssScrape() map[string]string {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)

	// 1. xml 을 struct 에다 env로 넣어두고... 거기서 불러와서
	// 2. 크롤링 하기 Title 하고 URL이면 될듯

	return nil

}

// OKKY 기술 게시글 읽기
func OkkyScrape() map[string]string {
	doc, err := goquery.NewDocument("https://okky.kr/")
	if err != nil {
		log.Fatal(err)
	}

	okkylist := make(map[string]string)

	// .Find(".class tag .class")
	// .Find("a").Text() a 태그 안에 든 거

	doc.Find(".article-middle-block").Each(func(i int, s *goquery.Selection) {

		title := s.Find("h5").Text()
		url := "https://okky.kr" + s.Find("h5 a").AttrOr("href", "없음")
		//okkylist = append(okkylist, title, url)

		okkylist[title] = url

	})

	return okkylist

}

// 깃허브 고 오픈소스 찾기
func GoScrape() map[string]string {

	// 깃허브에 연ㅇ결

	doc, err := goquery.NewDocument("https://github.com/trending/go?since=daily")
	if err != nil {
		log.Fatal(err)
	}

	githublist := make(map[string]string)

	// EachWithBreak 메서드의 true/false 리턴으로 검색 중 멈추고 싶을 때 멈출 수 있습니다.
	// 나는 너무많으니까 5개만 가져오는데서 멈춤

	var forLoop int = 0

	doc.Find(".repo-list li").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if forLoop > 4 {
			return false
		} else {
			// Trim 시리즈 적용하는 모듈 나중에 만들어야겠음(너무도 귀찮음)
			title := s.Find("h3 a").AttrOr("href", "없음")
			strings.TrimSpace(title)
			desc := s.Find(".py-1 p").Text()
			strings.TrimSpace(desc)
			strings.TrimLeft(desc, " ")

			githublist[title] = desc
			forLoop++
			return true
		}
	})

	return githublist
}

// IT 뉴스 찾기
func NewsScrape() map[string]string {

	doc, err := goquery.NewDocument("http://www.itworld.co.kr/")

	if err != nil {
		log.Fatal(err)
	}

	newslist := make(map[string]string)

	var forLoop int = 0

	doc.Find(".cio_summary").EachWithBreak(func(i int, s *goquery.Selection) bool {

		if forLoop > 4 {
			return false
		} else {
			title := s.Find("ul li a").Text()
			strings.TrimSpace(title)
			url := s.Find("ul li a").AttrOr("href", "없음")
			strings.TrimSpace(url)
			strings.TrimLeft(url, " ")

			newslist[title] = url
			forLoop++
			return true
		}
	})
	return newslist
}

// 문자열 슬라이싱용(귀찮음)
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
