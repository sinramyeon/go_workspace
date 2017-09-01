package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 그냥 인쇄
	fmt.Println(os.Args)

	title := flag.String("title", "", "영화이름") // 명령줄 옵션을 받아서 저장 (가운데는 베이직.)
	runtime := flag.Int("runtime", 0, "상영시간")

	flag.Parse()

	if flag.NFlag == 0 {
		// 명령줄 옵션 개수 0일시
		flag.Usage()
		return
	}

	fmt.Printf(
		"영화이름 : %s/n상영시간 : %d분",
		*title,
		*runtime, // 역참조
	)

}
