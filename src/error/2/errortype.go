package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time
	value int // 에러발생값
}

func (e HelloOneError) Error() string {
	return fmt.Sprintf("%v : %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	//의도적 에러
	if n == 1 {
		s := fmt.Sprintf("Hello", n)
		return s, nil
	}
	return "", HelloOneError{time.Now(), n} // 1이 아니면 에러. 에러 구조체를 생성해 리턴
}

func main() {
	s, err := helloOne(2) // 에러 발생
	if err != nil {
		log.Fatal(err) // 출력후 프로그램 종료
		// 종료하지 않고 에러 발생시키려면 log.Panic 또는 panic
		// 별 에러 아니면 log.Print도 괜찮음
	}
	fmt.Println(s) // 실행이 안됨
}
