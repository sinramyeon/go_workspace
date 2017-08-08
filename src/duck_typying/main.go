package main

import "fmt"

type Duck struct {
	//오리 구조체
}

func (d Duck) quack() {
	fmt.Println("꽥")
}

func (d Duck) feathers() {
	fmt.Println("오리 too much kawaii...")
}

type Person struct {
	// 사람 구조체
}

func (p Person) quack() {
	fmt.Println("꽥(죽음)")
}

func (p Person) feathers() {
	fmt.Println("사람 too less kawaii....")
}

type Quacker interface {
	quack()
	feathers() // 두 메서드를 갖는 인터페이스를 만들었다.
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck
	var trump Person // 분명 다른 타입의 인스턴스인데

	inTheForest(donald)
	inTheForest(trump)
	// 이러헥 둘 다 호출이 가능하다.
}
