package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/bunsenapp/go-selenium"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://ione.interpark.com/")

	println(doc.Text())

	if err != nil {
		log.Fatal(err)
	}

	// .Find(".class tag .class")
	// .Find("a").Text() a 태그 안에 든 거

	doc.Find(".innertabs").Each(func(i int, s *goquery.Selection) {

		title := s.Find("table").Text()

		url := s.Find("tr").AddClass(".viw-portlet-body")

		println(title)
		println(url)

	})
}

func main() {

	// Create a capabilities object.
	capabilities := goselenium.Capabilities{}

	// Populate it with the browser you wish to use.
	capabilities.SetBrowser(goselenium.FirefoxBrowser())

	// Initialise a new web driver.
	driver, err := goselenium.NewSeleniumWebDriver("http://localhost:9515/wd/hub", capabilities)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a session.
	_, err = driver.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer the deletion of the session.
	defer driver.DeleteSession()

	// Navigate to Google.
	_, err = driver.Go("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}

	// Hooray, we navigated to Google!
	fmt.Println("Successfully navigated to Google!")

	ExampleScrape()
}
