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
