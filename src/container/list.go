package main

import (
	"container/list"
	"fmt"
)

func main() {

	//Go 언어에서 제공하는 자료구조
	/*
	* 연결 리스트 : 각 노드를 한 줄로 연결한 자료구조
	* 힙 : 이진 트리를 활용한 자료구조
	* 링 : 각 노드가 원형으로 연결된 자료구조
	 */

	l := list.New() // 링크드 리스트 생성
	l.PushBack(10)  // 데이터추가
	l.PushBack(20)

	fmt.Println(l.Front().Value()) // 맨 앞 데이터. 맨 뒤는 Back()

	// Go언어 리스트는 이중 연결 리스트라서 양방향 순회 가능
	

}
