package test

import (
	"testing"
	"demo"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 测试失败时详细报告错误的行数
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := demo.Hello("Tom")
		want := "Hello Tom"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to everyone", func(t *testing.T) {
		got := demo.Hello("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}
