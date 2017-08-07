package main

import (
	"net"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
)

func main() {

	app := siris.New()

	app.Get("/", func(ctx context.Context) {

		ctx.Writef("ㅅㅓ ㅂ ㅓ")

	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("url path 연결 :  %s", ctx.Path())
	})

	// 사용자 tcp 리스너 등등 아무거나 연결 할 수 있어요
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}

	// 사용자 지정 리스너를 실행
	app.Run(siris.Listener(l))

}
