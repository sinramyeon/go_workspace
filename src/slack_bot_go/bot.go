// 봇 만들기

package main

import (
	_ "encoding/csv"
	"fmt"
	"log"
	_ "net/http"
	"strings"
)

func main() {

	ws, id := slackConnect("xoxb-220418545797-EE1cfi1N1E7CMaHoWBi8pP9K")
	fmt.Println("bot을 가동 합니다. 종료하려면 Ctrl + C를 누르세요.")

	for {

		// 입력하는 메시지 읽기
		m, err := getMessage(ws)

		if err != nil {
			log.Fatal("[ERROR] : ", err)
		}

		// 기사, 트렌드, 소식

		if m.Type == "message" && (m.Text == "트렌드" || m.Text == "소식" || m.Text == "기사") {

			fmt.Println("기사 크롤링 해오기")

			m.Text = "현재 " + m.Text + " 크롤링 중....\n" + "http://www.naver.com"
			//m.Attachments = append(m.Attachments, `{
			//"color": "#36a64f",
			//"text" : "Test"
			//}`)

			postMessage(ws, m)

		}

		// 오키, OKKY

		if m.Type == "message" && (m.Text == "오키" || strings.ToLower(m.Text) == "okky" || m.Text == "옼희" || m.Text == "오키") {

			fmt.Println("오키 구경하기")

			m.Text = "현재 " + m.Text + " 크롤링 중....\n" + "http://www.naver.com"

			postMessage(ws, m, OkkyScrape())
			//postMessage(ws, m)

		}

		// 깃허브, 깃헙, 오픈소스

		if m.Type == "message" && (m.Text == "깃허브" || m.Text == "깃헙" || m.Text == "오픈소스" || strings.ToLower(m.Text) == "github") {

			fmt.Println("GITHUB TRENDING NOW")

			m.Text = "현재 " + m.Text + " 크롤링 중....\n" + "http://www.naver.com"
			postMessage(ws, m)

		}

		// 트위터, 트윗, twitter, TWITTER

		if m.Type == "message" && (strings.ToLower(m.Text) == "twitter" || m.Text == "트윗" || m.Text == "트위터") {

			fmt.Println("SNS 구경하기")

			m.Text = "현재 " + m.Text + " 크롤링 중....\n" + "http://www.naver.com"
			postMessage(ws, m)

		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			parts := strings.Fields(m.Text)
			if len(parts) == 3 && parts[1] == "stock" {
				// looks good, get the quote and reply with the result
				go func(m Message) {
					m.Text = getQuote(parts[2])
					postMessage(ws, m)
				}(m)
				// NOTE: the Message object is copied, this is intentional
			} else {
				// huh?
				m.Text = fmt.Sprintf("sorry, that does not compute\n")
				postMessage(ws, m)
			}
		}

	}

}
