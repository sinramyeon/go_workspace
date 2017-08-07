package main

import (
	"net/url"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
	"github.com/go-siris/siris/core/host"
)

func main() {

	/*
		암호화 보안 프로토콜 : SSL과 TLS
		암호화 프로토콜은 보안 연결을 제공하여 두 당사자가 사생활 보호와 데이터 무결성을 갖고 서로 통신할 수 있돍 합니다.

		SSL의 발전형이 TLS입니다.append
		기밀성, 데이터 무결성, ID 및 디지털 인증서를 사용한 인증을 제공합니다.
		SSL 3.0과 TLS의 다양한 버전이 상호 운용되지 않으니 차이를 잘 봐야 합니다....
	*/

	app := siris.New()

	app.Get("/", func(ctx context.Context) {

		ctx.Writef("----SECURE----")

	})

	app.Get("/mypath", func(ctx context.Context) {

		ctx.Writef("/mypath 접속 시 SECURE서버")

	})

	// 8080 서버로 보다가 보안서버로 리다이렉트 될 시

	target, _ := url.Parse("https://127.0.1:443")
	go host.NewProxy("127.0.0.1:80", target).ListenAndServe()

	// 443포트로 HTTPS 서버 시작
	app.Run(siris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))
}
