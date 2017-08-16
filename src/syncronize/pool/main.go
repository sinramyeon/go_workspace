package main


import(
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct{
	tag string
	buffer []int
}

func main(){
	// 풀은 객체 사용 후 보관해두었다가 다시 사용하는 기능
	// 캐시라고 볼 수 있음
	
	runtime.GOMAXPROCS(runtime.NumCPU())

	pool := sync.Pool{			// 풀 할당
		New : Func() interface{} { // Get 함수 사용 시 호출될 함수 정의

			// 풀에 객체가 들어 있으면 New 필드 대신 보관된 객체 리턴

			data := new(Data) // 새 메모리 할당
			data.tag = "new"
			data.buffer = make([]int, 10)
			return data // 그 할당한 메모리 리턴
		},
	}

	for i:=0 ; i<10; i++{
		go func(){
			data := pool.Get().(*Data) // 풀에서 *Data타입으로 가져왔음
			// 풀에서 get 함수로 객체를 꺼낸 후엔 type assertion 필수
			//뭔소리야

			for index := range data.buffer{
				data.buffer[index] = ran.Intn(100)
			} // 슬라이스에 랜덤 값을 지정했다.

			fmt.Println(data)

			data.tag = "used"
			pool.Put(data) // 풀에 객체 보관
		}()
	}

	for i:=0 ; i<10 ; i++{
		go func(){
			data := pool.Get().(*Data)
			n:=0
			for index := range data.buffer{
				data.buffer[index] = n
				n+=2 // 슬라이스에 짝수를 넣었다.
			}

			fmt.Println(data)
			data.tag = "used"
			pool.Put(data) // 풀에 객체 보관
		}()
	}

	// sync.Pool 로 할당 후 Get, Put 함수로 사용
	// sync.Pool{ New : 어쩌구} 로 초기화 함수를 만듦
	

}