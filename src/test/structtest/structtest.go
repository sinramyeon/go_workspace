package main

import (
	"strconv"
)

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}

}
