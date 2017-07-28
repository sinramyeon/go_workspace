package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	println("Hello world")

	const (
		Apple = iota
		Grape
		Orange
	)

	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)

	println(f, u)

	var k int = 10
	var p = &k
	println(*p)

}
