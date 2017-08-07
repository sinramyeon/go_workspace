package main

//yaml은 xml, c, python, perl 에서 정의된 e mail 양식에서 개념을 얻어 만들어진 사람이 쉽게 읽을 수 있는 데이터 직렬화 양식입니다.
//xml, json파일같은거라고 생각하시면 됩니다.
//그리고 YAML 1.2기준으로 YAML파서는 JSON을 정상 파싱해 냅니다...

import (
	"github.com/go-siris/siris"
)

func main() {
	app := siris.New()
	app.Run(siris.Addr(":8080"), siris.WithConfiguration(siris.YAML("./configs/iris.yml")))

}

// 비슷한 방식으로 siris.TOML()로 .tml방식의 설정파일도 가져올 수 있습니다!
/*
.tml파일은 이렇게 생겼습니다.

DisablePathCorrection = false
EnablePathEscape = false
FireMethodNotAllowed = true
DisableBodyConsumptionOnUnmarshal = false
TimeFormat = "Mon, 01 Jan 2006 15:04:05 GMT"
Charset = "UTF-8"

[Other]
	MyServerName = "Siris"
*/
