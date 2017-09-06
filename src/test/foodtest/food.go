package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MeetUp struct {
	team string
	text string
	href string
	when string
	time string
}

func TrimTrim(s string) string {

	strings.Trim(s, " ")
	strings.TrimLeft(s, " ")
	strings.TrimPrefix(s, " ")
	strings.TrimRight(s, " ")
	strings.TrimSpace(s)

	return s

}
func main() {

	doc, _ := goquery.NewDocument("http://www.foodsafetykorea.go.kr/portal/healthyfoodlife/foodnutrient/simpleSearch.do?menu_no=403&menu_grp=MENU_GRP02&code4=2&code2=&search_name=%EB%82%99%EC%A7%80")
	th := doc.Find("div").AddClass("tab-pane active").Find("tr th a").Text()
	td := doc.Find("div").AddClass("tab-pane active").Find("td").Text()

	th = TrimTrim(th)
	td = TrimTrim(td)
	fmt.Println(th, td)

}
