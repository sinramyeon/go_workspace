package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true // 채널에 true 를 보내고 버퍼가 가득 차면 대기
			fmt.Println(i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // 버퍼에 값이 없으면 대기. 값이 있으면 꺼낸다.
		fmt.Println("메인함수 ", i)
	}
}

// 비동기 채널이므로 버퍼가 가득 찰 때까지 값을 계속 보냄
// 지금은 버퍼가 2개니까
// done 에 true 보내고 그 다음 루프에서 대기

// 01 메인 01 23 메인 23 식으로 출력(버퍼 가 2 개!!!)
// 비동기 채널의 구현과 동작
//보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고
//대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기
