package main

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/core/nettools"
)

func main() {
	app := siris.New()

	l, err := nettools.UNIX("/tmpl/srv.sock", 0666)
	// 연결할때 강제로 시도

	if err != nil {
		panic(err)
	}

	app.Run(siris.Listener(l))

}
