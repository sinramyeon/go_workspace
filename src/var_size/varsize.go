package main

import "fmt"
import "unsafe" //sizeof 함수

func main() {

	var num1 int8 = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int32 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Println(unsafe.Sizeof(num4))

}
