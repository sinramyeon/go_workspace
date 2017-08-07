// 웹에서 기사 가져오기

// 참고
//// rss
/*
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

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type slackReturn struct {
	Title string
	URL   string
}

// 트위터 읽기
func TwitterScrape(env twitterConfig) map[string]string {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	tweetlist := make(map[string]string)

	// https://apps.twitter.com/app/14097310/keys 를 보시면 됩니다.

	config := oauth1.NewConfig(env.confKey, env.confSecret)
	token := oauth1.NewToken(env.tokenKey, env.tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// 테스트 하시려면 유저 확인을 해보시면 좋습니다.

	//verifyParams := &twitter.AccountVerifyParams{
	//	SkipStatus:   twitter.Bool(true),
	//	IncludeEmail: twitter.Bool(true),
	//}

	//로그인 유저 얻기
	//user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	//fmt.Println("연결완료, %s", user)

	// 특정 계정 크롤링 하기

	tweets := GetUserTweets("Daeng_Daeng_yi", client)

	for _, v := range tweets {
		tweetlist[v.Text] = tweetlist[v.CreatedAt]
	}

	fmt.Println(tweetlist)

	return tweetlist

}

// 유저별 트윗을 얻는 모듈
// 얻어오고 싶은 양, 유저 아이디, 클라이언트 접속정보를 넣어 보세요...
func GetUserTweets(id string, client *twitter.Client) []twitter.Tweet {

	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      id,
		IncludeRetweets: twitter.Bool(false),
	})

	return tweets

}
