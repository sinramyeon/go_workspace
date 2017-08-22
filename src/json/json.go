package main

type Author struct {
	Name string `json:"name"`

	// json 문서 안 키를 소문자로 하려면 위와 같이 구조체 필드에 패그를 지정한다

}

func main() {

	// json.Unmarshal(읽을 내용, 저장할 곳 &포인터)

	// 보통 맵 map[string]interface{} 이나
	// struct 에 넣음

	// json.Marshal()로 값을 json문서로 변환 가능

}
