// slack 과 연결하기

// 참고

//// 고 슬랙봇
/*
https://github.com/BeepBoopHQ/go-slackbot

https://github.com/rapidloop/mybot  <--- 여기서는 요걸 이용
https://www.opsdash.com/blog/slack-bot-in-golang.html

https://github.com/nlopes/slack
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/websocket"
)

// struct
type responseRtmStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type responseSelf struct {
	Id string `json:"id"`
}

// websocket url과 id 반환

func slackStart(token string) (wsurl, id string, err error) {
	url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	var respObj responseRtmStart
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return
	}

	if !respObj.Ok {
		err = fmt.Errorf("Slack error: %s", respObj.Error)
		return
	}

	wsurl = respObj.Url
	id = respObj.Self.Id
	return
}

// slack 에서 메시지를 받아올 때 사용
type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	//Attachments []string `json:"attachments"`
}

func getMessage(ws *websocket.Conn) (m Message, err error) {
	err = websocket.JSON.Receive(ws, &m)
	return
}

// slack 에 메시지 전송
var counter uint64

/*
func postMessage(ws *websocket.Conn, m Message) error {

	fmt.Println(m)

	m.Id = atomic.AddUint64(&counter, 1)
	return websocket.JSON.Send(ws, m)
}
*/
/*
func postMessage(ws *websocket.Conn, m Message, s []string) error {

	m.Id = atomic.AddUint64(&counter, 1)
	m.

	return websocket.JSON.Send(ws, m)
}
*/

// gorequest 이용

func useRequest(m Message) {

	request := gorequest.New()
	// websocket 대신 gorequest로 바꿔써야지..
	resp, body, errs := request.Post("주소..").
		Set("파라미터", "값").
		Send(`{"JSON":"값", "json":"값"}`).
		End()

}

// 슬랙 연결
func slackConnect(token string) (*websocket.Conn, string) {
	wsurl, id, err := slackStart(token)
	if err != nil {
		log.Fatal(err)
	}

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		log.Fatal(err)
	}

	return ws, id
}
