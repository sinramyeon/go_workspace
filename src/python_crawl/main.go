// go slack API를 발견...ㅠㅠ
// https://github.com/nlopes/slack
// 위를 이용한 예제
// https://github.com/tcnksm/go-slack-interactive

package main

import (
	"log"
	"net/http"
	"os"

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
	env = envsetting(env)
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
	go slackListener.oscalScrape(tweetenv)

	//rtm := api.NewRTM()
	//go rtm.ManageConnection()

	log.Printf("[INFO] Server listening on :%s", env.Port)
	if err := http.ListenAndServe(":"+env.Port, nil); err != nil {
		log.Printf("[ERROR] %s", err)
		return 1
	}
	return 0

}

func (s *SlackListener) ListenAndResponse(tweetenv twitterConfig) {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

}

func (s *SlackListener) oscalScrape(tweetenv twitterConfig) {

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

		s.client.PostMessage("0scal", "", params)
	}
}
