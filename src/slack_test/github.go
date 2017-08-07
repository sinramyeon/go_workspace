package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	Pushevent        = "PushEvent"
	PullRequestEvent = "PullRequestEvent"
)

func getGitCommit(id string) bool {

	koryear, kormon, kordate := time.Now().Date()
	loc, _ := time.LoadLocation("Asia/Seoul")

	var commit_array []string

	// github 연결
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitconfig()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	events, _, err := client.Activity.ListEventsPerformedByUser(context.Background(), id, true, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range events {

		utctime := v.GetCreatedAt()
		year, month, day := utctime.In(loc).Date()

		// 오늘 한 커밋
		if ((v.GetType() == Pushevent) || (v.GetType() == PullRequestEvent)) && ((koryear == year) && (kormon == month) && (kordate == day)) {
			commitID := v.GetID()
			commit_array = append(commit_array, commitID)
		}
	}

	if len(commit_array) == 0 {
		return false
	} else {
		return true
	}

}
