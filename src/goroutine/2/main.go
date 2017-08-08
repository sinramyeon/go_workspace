package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)          // 랜덤한 숫자 100개 만듦
	time.Sleep(time.Duration(r)) // 그 랜덤한 숫자만큼 기다림
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // go루틴 100 개
	}

	fmt.Scanln()

	//고루틴을 종료하려면 함수 안에서 return 하거나 아니면 runtime.Goexit()함수를 활용 하세요
}
