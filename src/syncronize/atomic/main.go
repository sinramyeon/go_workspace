package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 원자적 연산 : 더 이상 쪼갤 수 없는 연산
	// 여러 스레드에서 같은 변수를 수정 시 서로 영향을 받지 않고 안전히 연산하도록 할 때 쓰임

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			data += 1 // 고루틴 2000개에서 data에 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			data -= 1 // 고루틴 1000개에서 data에서 1 씩 뺐음
			wg.Done()
		}()
	}

	wg.Wait()
	/// 결과는 2000-1000 = 1000이어야 하는데 값이 계속 바뀜
	// 뮤텍스 락 때 처럼 원자적 연산이 필요(ㄱ느데 뭔차이지??)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1) // 고루틴 2000개에서 원자적 연산으로 data에 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()

	//이렇게 하면 정확히 1000

}

/*
atomic 하고 mutex의 차이는 대체 뭐임 .... ?????

An atomic operation is one that cannot be subdivided into smaller parts. As such, it will never be halfway done, so you can guarantee that it will always be observed in a consistent state. For example, modern hardware implements atomic compare-and-swap operations.
atomic - 더이상 쪼갤 수 없는 단위
항상 같은 상태에서 고나측 될 것으로 기대할 수 있음

A mutex (short for mutual exclusion) excludes other processes or threads from executing the same section of code (the critical section). Basically, it ensures that at most one thread is executing a given section of code. A mutex is also called a lock.
mutex - 같은 상태의 코드에서 다른 스레드나 프로세스의 접근을 막음
한 스레드가 한 코드를 관리하게 해 줌
LOCK 이라고 생각하면 됨


atomic은 프로세서를 지원하고 전혀 락을 걸지 않음(os따라 달라짐). 대신 cpu 자원 확보는 안됨
mutex 락은 확실히 스레드를 일시 중단해서 cpu 자원을 확보할 수는 있는데 오버헤드를 발생시킬 수도 있음

*/
