package main

import (
	"github.com/go-siris/siris"
)

func main() {

	app := siris.New()

	app.Run(siris.Addr(":8080", siris.WithConfiguration(siris.Configuration{

		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 08 Aug 2017 15:04:05 GMT",
		Charset:                           "UTF-8",
	})))

	//또는 app.Run()전 app.Configure(siris.Withconf... )로 설정할수도 있다.

}
