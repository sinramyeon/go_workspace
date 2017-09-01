package calc

import "testing"

// 테스트는 항상 파일이름_test

func TestSum(t *testing.T) {
	//함수이름은 Test로 시작
	r := Sum(1, 2)
	if r != 3 {
		//기대하는 결과가 아닐 시 에러
		t.Error("뭐임")
	}
}

// 실행법
// 프로젝트 폴더로 가서 go test 하면 됨
