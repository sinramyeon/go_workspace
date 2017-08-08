package main

import(
	"fmt"
	"runtime"
)

// cpu 모든 코어를 사용하자(컴퓨터 불태움)

func main(){

	runtime.GOMAXPROCS(runtime.NumCPU()) // 최대 cpu 개수 설정
	s := "죽음이다..."

	for i= 0; i<100; i++{
		go func(n int){
			fmt.Println(s, n) // 익명함수 고루틴으로 실행
		}()
	}(i)

	fmt.Scanln()

}