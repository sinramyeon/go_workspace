// go slack API를 발견...ㅠㅠ
// https://github.com/nlopes/slack
// 위를 이용한 예제
// https://github.com/tcnksm/go-slack-interactive

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

type envConfig struct {
	// Port is server port to be listened.
	Port string `envconfig:"PORT" default:"3000"`

	// BotToken is bot user token to access to slack API.
	BotToken string `envconfig:"BOT_TOKEN" required:"true"`

	// VerificationToken is used to validate interactive messages from slack.
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`

	// BotID is bot user ID.
	BotID string `envconfig:"BOT_ID" required:"true"`

	// ChannelID is slack channel ID where bot is working.
	// Bot responses to the mention in this channel.
	ChannelID string `envconfig:"CHANNEL_ID" required:"true"`
}

type twitterConfig struct {
	confKey     string
	confSecret  string
	tokenKey    string
	tokenSecret string
}

type SlackListener struct {
	client    *slack.Client
	botID     string
	channelID string
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	// 유저 메시지 입력 외에도 이벤트 만들만한 것 생각해 보기...

	// 1. 설정
	var env envConfig
	env = envconfig(env)
	api := slack.New(env.BotToken)

	var tweetenv twitterConfig
	tweetenv = twitterconfig(tweetenv)

	slackListener := &SlackListener{
		client:    api,
		botID:     env.BotID,
		channelID: env.ChannelID,
	}

	// DEBUG설정 - 개발시에만 켜주세요
	//api.SetDebug(true)
	/* 로그인 테스트하기
	groups, err := api.GetGroups(false)

	if err != nil {
		fmt.Printf("%s 로그인 중 에러가 발생하였습니다. : %s\n", groups, err)
		return
	}
	*/

	// 2. 메시지 받기

	go slackListener.ListenAndResponse(tweetenv)

	//rtm := api.NewRTM()
	//go rtm.ManageConnection()

	http.Handle("/interaction", interactionHandler{
		verificationToken: env.VerificationToken,
	})

	log.Printf("[INFO] Server listening on :%s", env.Port)
	if err := http.ListenAndServe(":"+env.Port, nil); err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}
	return 0

	// 3. 봇 설정
	/*
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
	   				// 이 작업은 서버 연결 후 하는것이 좋겠군...

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

	   				if ev.User != info.User.ID && strings.Contains(ev.Text, "OKKY") || strings.Contains(ev.Text, "okky") || strings.Contains(ev.Text, "오키") || strings.Contains(ev.Text, "옼희") {

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

	   					m := RssScrape()

	   					for k, v := range m {

	   						attachment := slack.Attachment{

	   							Color: "#2a4f2e",
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

	   				if ev.User != info.User.ID && strings.Contains(ev.Text, "트윗") || strings.Contains(ev.Text, "트위터") || strings.Contains(ev.Text, "twitter") || strings.Contains(ev.Text, "TWITTER") {

	   					rtm.SendMessage(rtm.NewOutgoingMessage("트위터를 돌아보는 중입니다... :bird:", ev.Channel))

	   					m := TwitterScrape(tweetenv)

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
	   			//고통의 흔적들
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

	   			case *slack.InvalidAuthEvent:
	   				break Loop

	   			default:
	   				//Take no action
	   			}

	   		}
	   	}

	   }
	*/
}

func (s *SlackListener) ListenAndResponse(tweetenv twitterConfig) {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev, tweetenv); err != nil {
				log.Printf("[ERROR] 처리중 에러가 발생하였습니다.: %s", err)
			}
		}
	}
}

func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent, tweetenv twitterConfig) error {

	receivedMsg := ev.Msg.Text

	// 다른 채널에 쳤을때
	if ev.Channel != s.channelID {
		log.Printf("다른 채널 : %s %s", ev.Channel, s.channelID)
		return nil
	}

	// 봇에게 한 멘션이 아닐 때
	if !(strings.HasPrefix(receivedMsg, fmt.Sprintf("<@%s> ", "U6H0KD41X"))) {

		// 봇이 한 말이면 무시하자!
		if strings.Contains(ev.Msg.BotID, "bot") {
			return nil
		}

		if strings.Contains(receivedMsg, "기사") || strings.Contains(receivedMsg, "뉴스") || strings.Contains(receivedMsg, "소식") {

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

				s.client.PostMessage(ev.Channel, "", params)

			}
		}

		// 다. OKKY 입력 시

		if strings.Contains(receivedMsg, "오키") || strings.Contains(receivedMsg, "옼희") {

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

				s.client.PostMessage(ev.Channel, "", params)

			}

		}

		// 라. 블로그 입력 시(RSS)

		if strings.Contains(receivedMsg, "블로그") {

			m := RssScrape()

			for k, v := range m {

				attachment := slack.Attachment{

					Color: "#2a4f2e",
					Title: k,
					Text:  v,
				}

				params := slack.PostMessageParameters{

					Attachments: []slack.Attachment{
						attachment,
					},
				}

				s.client.PostMessage(ev.Channel, "", params)

			}

		}

		// 마. 트위터 입력 시

		if strings.Contains(receivedMsg, "트윗") || strings.Contains(receivedMsg, "트위터") {

			m := TwitterScrape(tweetenv)

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

				s.client.PostMessage(ev.Channel, "", params)
			}
		}

		// 바. 깃허브 입력 시(최신유행 GO 오픈소스 찾기)

		if strings.Contains(receivedMsg, "깃허브") || strings.Contains(receivedMsg, "깃헙") {

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

				s.client.PostMessage(ev.Channel, "", params)

			}

		}

		return nil
	}

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

	if _, _, err := s.client.PostMessage(ev.Channel, "", params); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}

	return nil
}
