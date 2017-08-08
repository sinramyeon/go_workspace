package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 클로저를 정의 후 고루틴으로 실행

	runtime.GOMAXPROCS(1) // CPU를 1개만 사용
	s := "인쇄~~~"

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n) // 익명함수를 고루틴으로 실행 == 클로저
		}(i)
	}

	//일반 클로저는 반복문 안에서 순서대로 실행되지만 고루틴으로 실행한 클로저는 반복문이 끝나고 고루틴 실행

	for i2 := 0; i2 < 100; i2++ {
		go func(n int) {
			fmt.Println(s, i2) // 익명함수를 고루틴으로 실행 == 클로저
		}() // 이렇게 실행하면
	}
	// 인쇄~~~ 100 이 100번 나옴(반복문이 다 끝나야 go func()가 실행 되니깐)

	fmt.Scanln()
}
