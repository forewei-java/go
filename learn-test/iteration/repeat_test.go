package iteration

import "testing"

/*
*
更多的 TDD 练习
学习了 for 循环
学习了如何编写基准测试 (go test -bench=.)
*/
func TestRepeat(t *testing.T) {
	assert := func(t *testing.T, repeated string, expected string) {
		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	}
	t.Run("test aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assert(t, repeated, expected)
	})

	t.Run("test aaaaaa", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assert(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
