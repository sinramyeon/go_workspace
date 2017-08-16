package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex) // 뮤텍스 조건변수 생성

	//대기중에는 고루틴 안 wait 함수, 대기하느 ㄴ고루틴을 깨울 때는 signal 함수
	//wait 부분은 mutex 보호

	c := make(chan bool, 3) //비동기 채널 생성

	for i := 3; i < 3; i++ {
		go func(n int) {
			mutex.Lock()
			c <- true // 잠궈놓고 c에 우선 true 를 보냈음
			fmt.Println(n)
			cond.Wait() //조건 변수 대기중......
			fmt.Println(n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c //채널에서 값을 꺼냄. 고루틴 3개 모두 실행되기 전 실행 안됨
	}

	for i := 0; i < 3; i++ {
		mutex.Lock()
		fmt.Println(i)
		cond.Signal() // 대기중인 고루틴을 하나씩 깨움
		mutex.Unlock()
	}

	cond.Broadcast() // 대기 중인 모든 고루틴을 깨울 수 있음.

	once := new(sync.Once) // 함수를 한 번만 실행할 Once

	for i := 0; i < 1000; i++ {
		go func(n int) {
			once.Do(hi)
		}(i)
	} //1000개의 고루틴 중 1개만 실행

}

func hi() {
	fmt.Println("1회")
}
