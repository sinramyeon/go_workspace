package main

func sum(a int, b int, c chan int) {
	c <- a + b //채널에 a랑 b 합 보내기
}

func main() {

	c := make(chan int) // int 형 채널
	go sum(1, 2, c)     // sum을 고루틴으로 실행 후 채널을 매개변수로 넘겨줘 봄

	n := <-c // 채널에서 값을 꺼내서 n에 넣어줘 봤음. n은 3이 됨

	// make chan 시 동기채널 생성
	// 채널을 매개변수로 받는 함수는 반드시 고루틴으로 실행

	//함수 예제
	// 변수명 chan 자료형 으로 매개변수로 넣을 수 있음
	// func sum(a int, b int, c chan int)

}
