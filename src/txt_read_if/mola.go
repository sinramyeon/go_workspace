package main

import "fmt"
import "io/ioutil"

func main() {

	var b []byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")
	//b에는 파일 내용, err에는 에러 값
	if err == nil {
		fmt.Println("으앙! %s", b)
	}

	// if 조건문 안에서 함수를 실행한 뒤 조건을 판단하고 싶을때는?

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Println("으앙! %s", b)
	} else {
		fmt.Println(err)
	}

	// 구문 밖에서는 변수를 사용할 수 없음

}
