package calc

import "testing"

type TestData struct {
	// 테스트 데이터 구조체
	arg1   int
	arg2   int
	result int
}

var testdata = []TestData{
	{2, 6, 8},
	{-8, 3, -5},
	{-6, -6, -12},
	{0, 0, 0},
}

func TestSum(t *testing.T) {
	for _, d := range testdata {
		r := Sum(d.arg1, d.arg2)
		if r != d.result {
			t.Errorf("에러 발생")
		}
	}
}
