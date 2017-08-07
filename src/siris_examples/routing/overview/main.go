package main

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
)

func main() {

	app := siris.New()

	app.Get("/", info)

	// GET: http://localhost:8080
	app.Get("/", info)

	// GET: http://localhost:8080/profile/kataras
	app.Get("/profile/{username:string}", info)
	// GET: http://localhost:8080/profile/kataras/backups/any/number/of/paths/here
	app.Get("/profile/{username:string}/backups/{filepath:path}", info)

	//아이콘
	//패비콘이란 인터넷 웹 브라우저의 주소창에 표시되는 웹사이트나 웹페이지를 대표하는 아이콘이다.
	// GET: http://localhost:8080/favicon.ico
	app.Favicon("./public/images/favicon.ico")

	// static 파일들

	// GET: http://localhost:8080/assets/css/bootstrap.min.css
	//	    참조하는 곳 ./public/assets/css/bootstrap.min.css
	// GET: http://localhost:8080/assets/js/react.min.js
	//      참조하는 곳 ./public/assets/js/react.min.js
	app.StaticWeb("/assets", "./public/assets")

	// 그룹짓기

	usersRoutes := app.Party("/users")

	// GET: http://localhost:8080/users/help
	usersRoutes.Get("/help", func(ctx context.Context) {

		ctx.Writef("GET, GET/$ID, POST, PUT/$ID, DELETE/$ID 를 해보세요 ")

	})

	usersRoutes.Get("/", func(ctx context.Context) {

		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("%s 님 로그인 하셨습니다.", id)

	})

	//app.Party
	// Party is just a group joiner of routes which have the same prefix and share same middleware(s) also.
	// Party could also be named as 'Join' or 'Node' or 'Group' , Party chosen because it is fun.
	// /users 로 시작하는 경우 -> app.Party("/users")

	// POST: http://localhost:8080/users
	usersRoutes.Post("/", func(ctx context.Context) {
		username, password := ctx.PostValue("username"), ctx.PostValue("password")
		ctx.Writef("유저 네임 = %s 비밀번호 = %s", username, password)
	})

	// PUT: http://localhost:8080/users
	usersRoutes.Put("/{id:int}", func(ctx context.Context) {
		id, _ := ctx.Params().GetInt("id") // ..Get("id")도 가능
		username := ctx.PostValue("username")
		ctx.Writef("update user for id= %d and new username= %s", id, username)
	})

	// DELETE: http://localhost:8080/users/42
	usersRoutes.Delete("/{id:int}", func(ctx context.Context) {
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("delete user by id: %d", id)
	})

	// 자기 호스트 설정을 손보세요.

	adminRoutes := app.Party("admin.")
	// GET: http://admin.localhost:8080
	adminRoutes.Get("/", info)
	// GET: http://admin.localhost:8080/settings
	adminRoutes.Get("/settings", info)

	// Wildcard 로 설정하고 싶으면
	dynamicSubdomainRoutes := app.Party("*.")

	// GET: http://아무거나.localhost:8080
	dynamicSubdomainRoutes.Get("/", info)

	if err := app.Run(siris.Addr(":8080")); err != nil {
		panic(err)
	}

}

func info(ctx context.Context) {
	method := ctx.Method()       // t서버 리소스 리퀘스트
	subdomain := ctx.Subdomain() // 서브도메인

	path := ctx.Path()

	paramsLen := ctx.Params().Len()

	ctx.Params().Visit(func(name string, value string) {
		ctx.Writef("%s = %s\n", name, value)
	})
	ctx.Writef("\nInfo\n\n")
	ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s\nParameters length: %d", method, subdomain, path, paramsLen)
}
