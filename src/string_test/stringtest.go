package main

import "fmt"
import "unicode/utf8"

func main() {

	var long_string = `여러 줄로 된
	문자열 입니다.`

	fmt.Println(len(long_string)) //인코딩 길이

	fmt.Println(utf8.RuneCountInString(long_string)) // 진짜길이

	// 문자열은 문자의 배열이므로 long_string[0] = '앵' 이런식으로
	// 내용을 직접 수정할 수는 없음

}
