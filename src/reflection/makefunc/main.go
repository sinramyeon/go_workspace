package main

import (
	"fmt"
	"reflect"
)

// 프로그래밍 언어에서 reflection이란 변수가 가진 데이터 타입을 인식해서 타입 캐스팅 등의 요건을 처리하는 구조
//

func h(args []reflect.Value) []reflect.Value {
	// 매개변수랑 리턴값은 무조건 []reflect.Value 로

	fmt.Println("함수")
	return nil
}

func main() {
	var hello func()                     // 함수를 담을 변수 선언
	fn := reflect.ValueOf(&hello).Elem() // hello 주소 값

	v := reflect.MakeFunc(fn.Type(), h) // h 함수 정보 생성

	fn.Set(v) // hello 값 정보인 fn에 h의 함수 정보 v를 설정해 함수를 연결

	hello()
}

// reflection의 동적 함수 생성 기능으로 타입별로 여러 번 함수를 구현하지 않아도 됨
