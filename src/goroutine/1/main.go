package main

import "fmt"

func hello() {
	fmt.Println("안녕")
}

func main() {
	go hello() // 스레드 생성보다 19234902490배 쉽고 리소스를 적게 사용
	// 동시에 함수를 실행 하므로 main, go 동시에 실행

	fmt.Scanln() // main 함수 종료되지 않도록 대기
}
