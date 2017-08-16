package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int `tag1: "이름" tag2 : "name"`
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var f floast62 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	t.Name()
	//자료형 이름
	t.Size()
	//자료형 크기
	t.Kind()
	//자료형 종류

	v.Type()
	//값이 담긴 변수 자료형 이름 등등등~~~~~

	data := Data{}

	a, b := reflect.TypeOf(data).FieldByName("a")
	a.Tag.Get("태그이름") // 태그 얻어오기
	//구조체 필드의 태그는 `태그명 : "내용"` 형식으로 저장

	var p *int = new(int) // 포인터 변수
	*p = 1

	//포인터 변수 값 가져오기

	reflect.ValueOf(p).Elem().Int()

	var i interface{}
	i = 1

	// 인터페이스 값 가져오기

	reflect.ValueOf(i).Elem().Int()
}
