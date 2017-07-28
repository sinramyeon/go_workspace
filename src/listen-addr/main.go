package main

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
)

func main() {
	app := siris.New()
	app.Get("/", func(ctx context.Context) {

		ctx.HTML("<h1>ㅇㅇ</h1>")
	})

	//에러 처리 하는 법?

	if err := app.Run(siris.Addr(":8080")); err != nil {
		panic(err)
	}
}
