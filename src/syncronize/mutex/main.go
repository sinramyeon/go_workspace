package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//mutex
	//상호 배제. 여러 스레드에서 공유되는 데이터를 보호

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 cpu 사용
	var data = []int{}                   //int형 슬라이스

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) //고루틴에서 1000번 반복해서 data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 cpu를 사용할 수 있도록 양보함
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) //고루틴에서 1000번 반복해서 data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 cpu를 사용할 수 있도록 양보함
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data)) // 2000번 추가했는데 1800~1990 사이 값 출력
	// 두 개 동시에 data 에 접근해서 append가 정확히 처리되지 않음

	// Race condition 경쟁 조건

	//-----------------

	// Mutex 사용시

	var data2 = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()            // 뮤텍스 잠금으로 보호
			data = append(data2, 1) //고루틴에서 1000번 반복해서 data 슬라이스에 1을 추가
			mutex.Unlock()
			runtime.Gosched() // 다른 고루틴이 cpu를 사용할 수 있도록 양보함
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()            // 뮤텍스 잠금으로 보호
			data = append(data2, 1) //고루틴에서 1000번 반복해서 data 슬라이스에 1을 추가
			mutex.Unlock()
			runtime.Gosched() // 다른 고루틴이 cpu를 사용할 수 있도록 양보함
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data2)) // 2000번 추가했고 2000 이 출력됨

	// mutex.Lock 과 Unlock은 항상 세트로! 짝이 맞지 않으면 Deadlock *교착상태 발생
}
