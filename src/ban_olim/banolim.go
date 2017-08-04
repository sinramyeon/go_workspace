package main

import "fmt"
import "math"

func main() {

	// 반올림 오차에 주의하자!
	// == 로 비교하면 안됨
	// 실수 비교시에는 머신...앱실론...(그게 뭐임)

	var a float64 = 3.24

	const epsilon = 1e-14

	fmt.Println(math.Abs(a-9.0) <= epsilon)

}
