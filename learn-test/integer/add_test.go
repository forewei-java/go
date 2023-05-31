package integer

import "testing"
import "fmt"


/**
测试驱动开发工作流的更多实践
整数，加法
编写更友好的文档，以便使用我们代码的人能够快速理解它的用法
如何使用我们代码的示例，这些代码作为测试的一部分被检查
**/
func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
