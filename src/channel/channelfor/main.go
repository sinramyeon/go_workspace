package main

import "fmt"

func producer(c chan<- int) {
	// 보내기 전용 채널

	//값의 흐름이 한 방향으로 고정

	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100 // 채널에 값 보냄

	// fmt.Println(<-c) 여기서 채널 값을 꺼낼 수 없음!
}

func consumer(c <-chan int) {
	//받기 전용 채널

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c) // 채널 값 꺼내기

	// c <- 1 채널에 값을 보낼 수 없음
}

func sum(a, b int) <-chan int {
	// 함수 리턴값이 int형 받기 전용 채널

	out := make(chan int)
	go func() {
		out <- a + b
	}() // 채널에 합 보냄

	return out
}

func num(a, b int) <-chan int {
	//int형 받기전용 채널을 리턴값으로

	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sum2(c <-chan int) <-chan int {
	//함수 매개변수로도 리턴값으로도 받기전용 채널을 써 봄

	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			//채널이 닫힐 때까지 값 꺼내기
			r = r + i //꺼낸 값 모두 더했음
		}
		out <- r //그 값을 방금 만든 채널에 넣어서
	}()

	return out
}

func main() {

	c := make(chan int)

	go producer(c)
	go consumer(c)
	// 결과 0 1 2 3 4 100 나옴

	c2 := sum(1, 2)

	fmt.Println(c2) //3

	c3 := num(1, 2)    // 1,2가 든 채널이 리턴되었습니다.
	out := sum2(c3)    // 그 값을 모두 더한 채널입니다
	fmt.Println(<-out) // 그 채널에 있는 값을 꺼냄

}

// send only chan <- int
// receive onlt <-chan int
