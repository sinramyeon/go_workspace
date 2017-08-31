package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	//의도적 에러
	if n == 1 {
		s := fmt.Sprintf("Hello", n)
		return s, nil
	}
	return "", fmt.Errorf("악", n) // 1이 아니면 에러
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
