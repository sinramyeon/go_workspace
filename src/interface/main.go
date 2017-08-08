package main

import "fmt"

type hello interface {
	//인터페이스를 정의한다
}

type myint int //int형을 정의하고

func (i myint) Print() {
	//myint랑 함수를 연결 했다.
	fmt.Println(i)
}

type Printer interface {
	Print() // Print 함수를 가지는 인터페이스를 정의
}

type Rectangle struct {
	//사각형 구조체
	width, height int
}

func (r Rectangle) Print() { // Rectangle 에 메서드 연결
	fmt.Println(r.height, r.width)
}

func main() {
	var h hello    //인터페이스 선언
	fmt.Println(h) //비어있으므로 nil 출력

	var i myint = 5
	var p Printer // 인터페이스 선언

	p = i     // p에 5 대입
	p.Print() // 인터페이스를 통해 5를 출력

	r := Rectangle{10, 20}

	var p2 Printer
	p2 = r
	p2.Print() // 여러 인터페이스 활용

	// Printer 안에 myint 자료형과 rectangle 구조체 둘다 들어있음

	pArr := []Printer{i, r} // 슬라이스 형태로 인터페이스 초기화
	for index, _ := range pArr {
		pArr[index].Print() // 슬라이스를 순회하면서 Print 를 호출
	}
}
