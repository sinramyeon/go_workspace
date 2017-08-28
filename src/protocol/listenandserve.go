package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux() // http 요청 멀티플렉서

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		res.Write([]byte("Hello")) // 브라우저 응답

	})

	http.ListenAndServe(":80", mux) // 80포트에 연결

}
