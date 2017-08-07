package main

import (
	"github.com/go-siris/siris"
)

func main() {
	app := siris.New()

	// 이런식으로 설정할수도 있으니 참고하자.
	app.Run(siris.Addr(":8080"), siris.WithoutBanner, siris.WithCharset("UTF-8"))

}
