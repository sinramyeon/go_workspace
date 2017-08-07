package main

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
	"github.com/go-siris/siris/view"
)

// User 설정
type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {

	app := siris.New()

	// html/template 엔진을 이용해서 템플릿 정의
	// ./views폴더 안 .html 파일을 모두 찾아옴
	// .Reload(true) 리퀘스트마다 템플릿 다시 로드
	app.AttachView(view.HTML("./views", ".html"))

	// 에러 핸들링

	app.OnErrorCode(siris.StatusInternalServerError, func(ctx context.Context) {

		// .Values() 로 핸들러와 미들웨어 간 통신

		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("서버 에러 발생 : %s", errMessage)
			return
		}

		ctx.Writef("알수없는 서버 에러 발생")

	})

	app.Use(func(ctx context.Context) {
		ctx.Application().Log("%s 에서 리퀘스트 시작", ctx.Path())
		ctx.Next()
	})

	//Post 로 들어온 내용
	// Method POST: http://localhost:8080/decode
	app.Post("/decode", func(ctx context.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s %s 는 %d살이며 %s에서 왔습니다.", user.Firstname, user.Lastname, user.Age, user.City)
	})

	// Method GET: http://localhost:8080/encode
	app.Get("/encode", func(ctx context.Context) {
		doe := User{
			Username:  "Johndoe",
			Firstname: "John",
			Lastname:  "Doe",
			City:      "Neither FBI knows!!!",
			Age:       25,
		}

		ctx.JSON(doe)
	})

	// Method GET: http://localhost:8080/profile/anytypeofstring
	app.Get("/profile/{username:string}", profileByUsername)

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		// Method GET: http://localhost:8080/users/42
		usersRoutes.Get("/{id:int min(1)}", getUserByID)
		// Method POST: http://localhost:8080/users/create
		usersRoutes.Post("/create", createUser)
	}

	// 들어오는 로컬호스트 8000의 HTTP/1.x & HTTP/2 clients 를 실행하고 있음
	app.Run(siris.Addr(":8080"), siris.WithCharset("UTF-8"))
}

func logThisMiddleware(ctx context.Context) {
	ctx.Application().Log("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	//.Next는 다음으로 이동하려고 씁니다.
	// 다음이 없으면 멈춥니다!
	ctx.Next()
}

func profileByUsername(ctx context.Context) {
	// context.params().get()으로 파라미터명에 따른 파라미터를 갖고 옵니다.
	username := ctx.Params().Get("username")
	ctx.ViewData("Username", username)
	// AttachView에 설정된 "./views/" 아래 "users/profile.html"
	// viewData("Username" 안에 username 값)을 담고, 사용은 {{ Username }}같이 하시면 됩니다.
	ctx.View("users/profile.html")
}

func getUserByID(ctx context.Context) {
	userID := ctx.Params().Get("id") // 또는 : .Values().GetInt/GetInt64 등으로도 활용 가능
	// DB연동 구문과 연결해도 되겠죠
	user := User{Username: "username" + userID}

	ctx.XML(user)
}

func createUser(ctx context.Context) {
	var user User
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "유저 생성 중 오류가 발생하였습니다. "+err.Error())
		ctx.StatusCode(siris.StatusInternalServerError)
		return
	}
	// "./views/users/create_verification.html"
	// {{ . }} 는 아래와 같은 User 오브젝트와 같음. : {{ .Username }} , {{ .Firstname}}
	ctx.ViewData("", user)
	ctx.View("users/create_verification.html")
}
