package tools

import (
	"fmt"
	"testing"
)

//基准测试
func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

func TestAdd(t *testing.T) {
	if Add(1, 2) == 3 {
		t.Log("1+2=3")
	}

	if Add(1, 1) == 3 {
		t.Error("1+1=3")
	}
}
