package main

import (
	"container/ring"
	"fmt"
)

func main() {

	data := []string{"Maria", "John", "Andrew"}

	r := ring.New(len(data)) // data길이의 링 생성
	for i := 0; i < r.Len(); i++ {
		r.Value = data[i]
		r = r.Next() // 값을 넣고 다음으로 이동
	}

	r.Do(func(x interface{}) {
		fmt.Println(x) // 순회
	})
}

// 이중으로 연결된 원 리스트
// 처음이 없고 nli도 없음
