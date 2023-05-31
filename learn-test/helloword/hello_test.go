package main

import "testing"

/**
编写一个失败的测试，并查看失败信息，我们知道现在有一个为需求编写的 相关 的测试，并且看到它产生了 易于理解的失败描述
编写最少量的代码使其通过，以获得可以运行的程序
然后 重构，基于我们测试的安全性，以确保我们拥有易于使用的精心编写的代码
**/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Elodie", "")
		want := "Hello, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
