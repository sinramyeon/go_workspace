// go slack API를 발견...ㅠㅠ
// https://github.com/nlopes/slack
// 위를 이용한 예제
// https://github.com/tcnksm/go-slack-interactive

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

	// 연결 오류가 났을 때 panic / recover 용 defer함수 만들기
	// 기타 panic 날법해 보이는 애들 처리
	// 유저 메시지 입력 외에도 이벤트 만들만한 것 생각해 보기...

	// 1. 설정
	var env envConfig
	env = envconfig(env)
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
									{
										Text:  "BLOG",
										Value: "BLOG",
									},
									{
										Text:  "TWITTER",
										Value: "TWITTER",
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

				// https://ngrok.com/ 으로 연결하던 서버를 통해 slack에서 보내는 메시지를 받아와야 한다고 함
				// https://github.com/nlopes/slack/blob/master/attachments.go

				// 어디 기사사이트로 할지 아직 못정해서 안만듦.
				// 나. 기사, 뉴스, 소식 입력 시

				if ev.User != info.User.ID && strings.Contains(ev.Text, "기사") || strings.Contains(ev.Text, "뉴스") || strings.Contains(ev.Text, "소식") || strings.Contains(ev.Text, "NEWS") || strings.Contains(ev.Text, "news") {

					rtm.SendMessage(rtm.NewOutgoingMessage("신문의 IT 섹션을 펼치는 중... :camera_with_flash:", ev.Channel))

					m := NewsScrape()

					for k, v := range m {

						attachment := slack.Attachment{

							Color: "#cc1512",
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

				// 다. OKKY 입력 시

				if ev.User != info.User.ID && strings.Contains(ev.Text, "OKKY") || strings.Contains(ev.Text, "okky") || strings.Contains(ev.Text, "오키") {

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

				// 라. 블로그 입력 시(RSS)

				if ev.User != info.User.ID && strings.Contains(ev.Text, "blog") || strings.Contains(ev.Text, "BLOG") || strings.Contains(ev.Text, "블로그") {

					rtm.SendMessage(rtm.NewOutgoingMessage("기술 블로그 구경 중입니다... :red_car:", ev.Channel))

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

				// 마. 트위터 입력 시

				if ev.User != info.User.ID && strings.Contains(ev.Text, "트윗") || strings.Contains(ev.Text, "트위터") {

					rtm.SendMessage(rtm.NewOutgoingMessage("트위터를 돌아보는 중입니다... :bird:", ev.Channel))

					m := TwitterScrape()

					for k, v := range m {

						attachment := slack.Attachment{

							Color: "#42c7d6",
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

				// 바. 깃허브 입력 시(최신유행 GO 오픈소스 찾기)

				if ev.User != info.User.ID && strings.Contains(ev.Text, "github") || strings.Contains(ev.Text, "GITHUB") || strings.Contains(ev.Text, "깃허브") || strings.Contains(ev.Text, "깃헙") {

					rtm.SendMessage(rtm.NewOutgoingMessage("최신유행 GO 오픈소스를 공부하러 가는 중... :lollipop: ", ev.Channel))

					m := GoScrape()

					for k, v := range m {

						title := strings.TrimPrefix(k, "/")
						title_link := "https://github.com" + strings.TrimSpace(k)

						attachment := slack.Attachment{

							Color:     "#f7b7ce",
							Title:     title,
							TitleLink: title_link,
							Text:      v,
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
			/*고통의 흔적들
			case *slack.AttachmentAction:
				fmt.Println("AttachmentAction")
			case *slack.AttachmentActionCallback:
				fmt.Println("AttachmentActionCallback")
			case *slack.AttachmentActionOption:
				fmt.Println("AttachmentActionOption")
			case *slack.AttachmentActionOptionGroup:
				fmt.Println("AttachmentActionOptionGroup")
			case *slack.AttachmentField:
				fmt.Println("엥")
			*/

			case *slack.InvalidAuthEvent:
				break Loop

			default:
				//Take no action
			}

		}
	}

}
