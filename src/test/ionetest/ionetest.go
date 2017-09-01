package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

type Position struct {
	KR string
	US string
}

type Authorname struct {
	KR string
	US string
}

type Notice struct {
	Day        string
	AuthorName string
	Position   string
	Text       string
}

type EntryData struct {
	Key   string `xml:"name,attr"`
	Value string `xml:"text"`
}

type ViewEntry struct {
	Key   string      `xml:"unid,attr"`
	Value []EntryData `xml:"entrydata"`
}
type ViewEntries struct {
	XMLName     xml.Name    `xml:viewentries`
	ViewEntries []ViewEntry `xml:"viewentry"`
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	parsed := new(ViewEntries)
	_, body, _ := gorequest.New().Get(
		"http://ione.interpark.com/gw/app/bult/bbslink.nsf/wviwportalnotice?ReadViewEntries&restricttocategory=01&start=1&count=10&page=1",
	).Type("xml").AddCookie(
		&http.Cookie{Name: "LtpaToken", Value: "AAECAzU5QTY3Njg4NTlBN0M4MDhDTj0Ruc4RvcIRseIvT1U9MjAxMTAzMzQvTz1pbnRlcnBhcms/jQWHU+jsjHSmyRqoj3Goj/z8Qg=="},
	).End()

	_ = xml.Unmarshal([]byte(body), &parsed)

	var notice Notice
	var noticelist []Notice

	for _, v := range parsed.ViewEntries {
		var entrydata []EntryData
		entrydata = v.Value
		for key, val := range entrydata {

			if notice.AuthorName != "" && notice.Day != "" && notice.Text != "" && notice.Position != "" {
				noticelist = append(noticelist, notice)
				notice.AuthorName = ""
				notice.Day = ""
				notice.Text = ""
				notice.Position = ""
			}

			switch key {
			case 0:
				notice.Day = val.Value
			case 1:
				notice.Text = val.Value
			case 2:
				notice.AuthorName = val.Value
			case 3:
				notice.Position = val.Value
			}
		}
	}

	fmt.Println(noticelist)

}
