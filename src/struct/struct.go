package main

import "fmt"



	type Person struct {
		name string
		age int
	}

	func (p *Person) greeting(){
		fmt.Println("씨발 왜안나와")
	}
	type Student struct{
		p Person // 학생 구조체 안 사람 구조체 필드
		school string
	}


func main(){
	var s Student
	s.p.greeting()

}