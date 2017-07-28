package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4

	max_xy := max(x, y)
	print(max_xy)

	r := sum(1, 2)
	fmt.Println(r)

}

func sum(a int, b int) (r int) {
	r = a + b // 리턴값 변수 r에 값 대입
	return    // 리턴값 변수를 사용할 때는 return 뒤에 변수를 지정하지 않음
}

// ´_`??
