// go slack API를 발견...ㅠㅠ
// https://github.com/nlopes/slack

package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

type interactionHandler struct {
	slackClient       *slack.Client
	verificationToken string
}

type envConfig struct {
	botToken  string
	botId     string
	channelId string
}

func main() {

	// 1. 설정
	var env envConfig
	env = envconfig(env)

	fmt.Println(env)

	api := slack.New(env.botToken)
	// DEBUG설정 - 개발시에만 켜주세요
	//api.SetDebug(true)

	groups, err := api.GetGroups(false)

	if err != nil {
		fmt.Printf("%s 로그인 중 에러가 발생하였습니다. : %s\n", groups, err)
		return
	}

	// 2. 메시지 받기
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// 3. 봇 설정

Loop:
	for {

		select {

		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {

			// 연결 case *slack.ConnectedEvent:

			// 메시지 수신
			case *slack.MessageEvent:

				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				// 가. 봇에게 멘션 시

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {

					attachment := slack.Attachment{

						Text:       "오늘의 핫한 소식들 듣고 가실래요? :newspaper: ",
						Color:      "#f9a41b",
						CallbackID: "news",
						Actions: []slack.AttachmentAction{

							{

								Name: actionSelect,
								Type: "select",

								Options: []slack.AttachmentActionOption{

									{
										Text:  "IT News",
										Value: "IT News",
									},
									{
										Text:  "OKKY",
										Value: "OKKY",
									},
									{
										Text:  "Go opensource",
										Value: "Go opensource",
									},
								},
							},
						},
					}

					params := slack.PostMessageParameters{

						Attachments: []slack.Attachment{
							attachment,
						},
					}

					rtm.PostMessage(env.channelId, "", params)

				}

				// 나. 기사, 뉴스, 소식 입력 시

				if strings.Contains(ev.Text, "기사") || strings.Contains(ev.Text, "뉴스") || strings.Contains(ev.Text, "소식") || strings.Contains(ev.Text, "NEWS") || strings.Contains(ev.Text, "news") {

					attachment := slack.Attachment{

						Pretext: "IT 뉴스 :camera_with_flash:  ",
						Color:   "#104293",
						Title:   "기사 제목",
						//TitleLink: "http://www.naver.com",
						Text: "http://www.naver.com",
					}

					params := slack.PostMessageParameters{

						Attachments: []slack.Attachment{
							attachment,
						},
					}

					rtm.PostMessage(env.channelId, "", params)
				}

				// 다. OKKY 입력 시

				if strings.Contains(ev.Text, "OKKY") || strings.Contains(ev.Text, "okky") || strings.Contains(ev.Text, "오키") {

					rtm.SendMessage(rtm.NewOutgoingMessage("okky 기술 글들을 긁어오는 중입니다... :desktop_computer:", ev.Channel))

					m := OkkyScrape()

					for k, v := range m {

						attachment := slack.Attachment{

							Color: "#104293",
							Title: k,
							Text:  v,
						}

						params := slack.PostMessageParameters{

							Attachments: []slack.Attachment{
								attachment,
							},
						}

						rtm.PostMessage(env.channelId, "", params)

					}

				}

			case *slack.RTMError:

			case *slack.InvalidAuthEvent:
				break Loop

			default:
				//Take no action
			}

		}
	}

}
