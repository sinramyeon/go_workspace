package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {

	data := make([]byte, 4096) // 슬라이스 만들고

	for {

		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n])) // 데이터 출력

		c.Write(data[:n]) //클라이언트로 데이터 전송
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {

	ln, err := net.Listen("tcp", ":8000") // 8000포트로 tcp 연결
	if err != nil {
		return
	}

	defer ln.Close() // 항상 연결은 끝나기 전 닫아줘 버릇합시다

	for {

		conn, err := ln.Accept() // 연결 시 cep연결을 리턴
		if err != nil {
			return
		}
		defer conn.Close()

		go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 돌림

	}

}
