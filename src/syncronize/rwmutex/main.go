package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int = 0

	go func() {
		for i := 0; i < 3; i++ {
			data += 1
			fmt.Println(data)
			time.Sleep(10 * time.Millisecond) // 데이터에 값 쓰고 10밀리초 대기
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(data) // 데이터 값을 읽기
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(data) // 데이터 값을 읽기
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)

	//실행결과 111 2 333333 식으로 읽다 쓰다가 반복되서 엉망이 됨
	//읽기 락 Read Lock(쓰기 락 방지)과 쓰기 락 Write Lock(읽기 쓰기 방지) 로 보장성을 높여 보자.

	var rwMutex = new(sync.RWMutex) // 읽기쓰기 뮤텍스

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 뮤텍스를 잠궈 쓰기 보호 시작
			data += 1
			fmt.Println(data)
			time.Sleep(10 * time.Millisecond) // 데이터에 값 쓰고 10밀리초 대기
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()   // 읽기 뮤텍스 잠금으로 읽기 보호 시작
			fmt.Println(data) // 데이터 값을 읽기
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()   // 읽기 뮤텍스 잠금으로 읽기 보호 시작
			fmt.Println(data) // 데이터 값을 읽기
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	//실행결과 001112223 으로 읽기 동작이 모두 끝난 후 쓰기 동작이 시행....
	//읽기 동작끼리는 서로를 막지 않으나 쓰기 읽기 간 막음
	//읽기 쓰기 뮤텍스로 다른곳에서 이전 데이터를 읽지 못하도록 방지하거나 읽기 중 데이터가 바뀌려 하는 것을 방지
	//읽기 동작이 많을 때 유리

}
