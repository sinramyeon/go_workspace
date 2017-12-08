package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now().Local()
	dueDate := time.Date(int(2017), time.December, int(16), int(0), int(0), int(0), int(0), time.UTC)
	diff := timeNow.Sub(dueDate)

	d_day := int(diff.Hours() / 24)
	fmt.Println(d_day)
	fmt.Println(timeNow)
	fmt.Println(dueDate)

}
