package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

// 구조체 집합
type Students []Student

func (s Students) Len() int {
	return len(s) // 데이터의 길이를 구함. 슬라이스이므로 len 함수를 사용
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name // 학생 이름순으로 정렬
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // 두 데이터의 위치를 바꿈
}

type ByScore struct { // 점수순 정렬을 위해 구조체 정의
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score // 학생 점수순으로 정렬
}

////////------------

type By func(s1, s2 *Student) bool // 각 상황별 정렬 함수를 저장할 타입

func (by By) Sort(students []Student) { // By 함수 타입에 메서드를 붙임
	sorter := &studentSorter{ // 데이터와 정렬 함수로 sorter 초기화
		students: students,
		by:       by,
	}
	sort.Sort(sorter) // 정렬
}

type studentSorter struct {
	students []Student                  // 데이터
	by       func(s1, s2 *Student) bool // 각 상황별 정렬 함수
}

func (s *studentSorter) Len() int {
	return len(s.students) // 슬라이스의 길이를 구함
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j]) // by 함수로 대소관계 판단
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i] // 데이터 위치를 바꿈
}

func main() {
	s := []Student{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	name := func(p1, p2 *Student) bool { // 이름 오름차순 정렬 함수
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool { // 점수 오름차순 정렬 함수
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool { // 점수 내림차순 정렬 함수
		return !score(p1, p2)
	}

	By(name).Sort(s) // name 함수를 사용하여 이름 오름차순 정렬
	fmt.Println(s)

	By(reverseScore).Sort(s) // reverseScore 함수를 사용하여 점수 내림차순 정렬
	fmt.Println(s)

	s2 := Students{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	sort.Sort(s2) // 이름을 기준으로 오름차순 정렬
	fmt.Println(s2)

	sort.Sort(sort.Reverse(ByScore{s})) // 점수를 기존으로 내림차순 정렬
	fmt.Println(s)

	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Float64Slice(b))
	sort.Sort(sort.StringSlice(c))

	fmt.Println(a, b, c)

}
