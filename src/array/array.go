package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for x := 0; x < len(a); x++ {
		fmt.Println(a[x])
	}

	// 배열 길이 구하지 않고 가져오기
	// for 인덱스, 값 = range 배열{}

	for i, value := range a {
		fmt.Println(i, value)
	}
}
