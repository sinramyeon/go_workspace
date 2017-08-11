package main

import (
	"fmt"
	"time"
)

func main() {
	donggi := make(chan bool) //동기화 채널 생성
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			donggi <- true              // 고루틴엔 우선 true값을 보냈음
			fmt.Println("고루틴 : ", i)    // 반복문의 변수 출력
			time.Sleep(1 * time.Second) // 1초 대기

			//동기 채널 donggi : 다른 쪽에서 값 꺼낼 때까지 대기
			//메인 함수 반복문에서 donggi 값 꺼낸 후 다시 이 for문이 돌아감
		}
	}() // 익명함수로 실행

	for i := 0; i < count; i++ {
		<-donggi // 채널에 값 들어올때까지 대기, 값 꺼냄
		fmt.Println(i)
	}

}
