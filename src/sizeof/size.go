package main

import "fmt"
import "unsafe"

func main() {

	var num1 int8 = 1
	fmt.Println(unsafe.Sizeof(num1))

	const age int = 10

	// age = 30 컴파일에러 발생

	const (
		a = iota
		b
		c
		d
		e
		f
	)

	fmt.Println(a)
	fmt.Println(f)

	over := 1
	print(over)

	for i := 5; i > 0; i-- {

		print(i)

	}

	myappend := []int{1, 2, 3}
	toappend := []int{4, 5, 6}

	myappend = append(myappend, toappend...)

	foo(1, 2)
	foo(1, 2, 3)

	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func foo(args ...int) {
	for i, a := range args {
		fmt.Printf("%v. %v\n", i, a)
	}
}
