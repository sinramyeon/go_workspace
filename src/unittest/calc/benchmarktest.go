package calc

import "testing"

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

// 결괏값 검증 및 성능도 검사
// go test -bench 로 실행하면 됨
