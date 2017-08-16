package main

import (
	"fmt"
	"time"
)

func main() {
	//여러 채널을 손쉽게 쓰려면 select
	/*
		select{
			case <- 채널1 :
			case <- 채널2...
			default :
				// 모든 채널에 값이 들어오지 않았을 때
		}
	*/

	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}() // c1에 10 보내고 100밀리초 대기

	go func() {
		for {
			c2 <- "메롱"
			time.Sleep(500 * time.Millisecond)
		}
	}() // c2에 메롱 보내고 500밀리초 대기

	go func() {
		for {
			select {
			case i := <-c1:
				// 채널 1 에 값이 들어왔을 경우
				fmt.Println(i)
			case s := <-c2:
				//패널 2에 들어왔을 경우
				fmt.Println(s)
			case <-time.After(50 * time.Microsecond):
				//50마이크로초 후 현재 시간이 담긴 채널이 리턴
				fmt.Println("타임아웃")
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초간 프로그램 실행
}
