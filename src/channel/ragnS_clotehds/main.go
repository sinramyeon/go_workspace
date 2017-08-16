package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			//0 1  2= 3 맞춰서
			c <- i // 채널에 값 값 보낸 후

		}
		close(c) // 채널 종료
	}()

	for i := range c { //채널 값 차례로 꺼내 출력
		fmt.Println(i)
	}

}

// range 와 close함수
/*
 이미 닫힌 채널에 값을 보낼 시 패닉 발생
 채널 닫을 시 range 루프 종료
  채널이 열려 있고 값이 들어오지 않을 시 range는 계속 대기
*/
