package main

import (
	"github.com/go-siris/siris/context"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/view"
)

// 쓸 구조체

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

// 메인 함수

func main() {
	app := siris.New()

	// view 추가 하기
	// views 폴더 아래 .html인거 찾음
	// 리퀘스트 마다 리로드 함(개발모드!!)
	app.AttachView(view.HTML("./views", ".html").Reload(true))

	// http Error handling
	app.OnErrorCode(siris.StatusInternalServerError, func(ctx context.Context) {

		// 핸들러랑 미들웨어 사이에서 .Values 이용
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {

			ctx.Writef("서버 에러가 발생했습니다. : %s", errMessage)
			return
		}

		ctx.Writef("서버 에러인데 뭔지 모르겠습니다.")

	})

	app.Use(func(ctx context.Context) {
		ctx.Application().Log("%s로 접속중...", ctx.Path())
		ctx.Next()
	})

	// POST http://localhost:8080/decode

	app.Post("/decode", func(ctx context.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s 는 %d 살이고 %s에서 왔어요~", user.Firstname, user.Lastname, user.Age, user.City)
	})

	// GET: http://localhost:8080/encode
	app.Get("/encode", func(ctx context.Context) {
		doe := User{
			Username:  "김설화",
			Firstname: "김",
			Lastname:  "설화",
			City:      "서울",
			Age:       26,
		}

		ctx.JSON(doe)
	})

	// GET http://localhost:8080/profle/anytypeofstring

	app.Get("/profile/{username:string}", profileByUsername)

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		// GET: http://localhost:8080/users/42
		usersRoutes.Get("/{id:int min(1)}", getUserByID)
		// POST: http://localhost:8080/users/create
		usersRoutes.Post("/create", createUser)
	}

	// 서버 시작

	app.Run(siris.Addr(":8080"), siris.WithCharset("UTF-8"))

}

// 함수들

func logThisMiddleware(ctx context.Context) {

	ctx.Application().Log("주소 : %s | IP : %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func profileByUsername(ctx context.Context) {
	// . 현재 위치
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)

	// django render 처럼 여기도 렌더가 있나 봄...
	// {{ .uUername }} 식으로 쓴다고 하네요. django와 같군..

	ctx.View("users/profile.html")
}

func getUserByID(ctx context.Context) {

	userID := ctx.Params().Get("id")
	//나중에 db연동하면 여기에서 select 해오거나 하면 되겠죠?
	user := User{Username: "username" + userID}

	ctx.XML(user)
}

func createUser(ctx context.Context) {
	var user User
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
		ctx.StatusCode(siris.StatusInternalServerError)
		return
	}
	// renders "./views/users/create_verification.html"
	// with {{ . }} equals to the User object, i.e {{ .Username }} , {{ .Firstname}} etc...
	ctx.ViewData("", user)
	ctx.View("users/create_verification.html")
}
