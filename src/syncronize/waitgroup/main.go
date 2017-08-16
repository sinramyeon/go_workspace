package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// 댜기 그룹은 고루틴이 모두 끝날 때까지 기다릴 때 사용
	// tim.Sleep 나 fmt.Sacnln 대신 얘를 써보자!

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup) // 대기 그룹

	for i := 0; i < 10; i++ {
		wg.Add(1) // 반복시마다 wg.Add로 1씩 추가
		go func(n int) {
			//고루틴 10개 만들어짐
			fmt.Println(n)
			wg.Done() // 고루틴 끝
		}(i)
	}

	wg.Wait() // 고루틴 다 끝날땎지 대기

	// * 주의 *
	// Add 함수에 설정한 값과 Dond 함수 호출되는 횟수는 같게
	//Add(3) 이면 Done()도 세번 호출

	//지연 호출로 defer wg.Done()으로 하 ㄹ수도 있음

}
