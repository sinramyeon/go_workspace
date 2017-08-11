package main

import (
	"fmt"
)

func sum(a, b int, c chan int) {
	c <- a + b // 채널에 a와 b의 합을 보냄
}

func main() {
	// make()로 공간 할당 - 동기 채널 생성 synchronous channel
	c := make(chan int)
	// int 형 채널

	go sum(1, 2, c) // 채널을 매개변수로 받는 함수는 반드시 go 키워드로 고루틴으로 실행

	n := <-c //채널 값을 n에 대입. 3
	fmt.Println(n)
}
