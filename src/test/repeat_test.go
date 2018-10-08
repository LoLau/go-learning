package test

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

//benchmark 基准测试
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Repeat("a", 5)
	}
}

func Repeat(s string, repeatCount int) string {
	var str string
	for i := 1; i <= repeatCount; i ++ {
		str += s
	}
	return str
}
