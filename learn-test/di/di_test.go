package main

import (
	"bytes"
	"testing"
)

/*
*
测试代码。如果你不能很轻松地测试函数，这通常是因为有依赖硬链接到了函数或全局状态。例如，如果某个服务层使用了全局的数据库连接池，这通常难以测试，并且运行速度会很慢。DI 提倡你注入一个数据库依赖（通过接口），然后就可以在测试中控制你的模拟数据了。
关注点分离，解耦了数据到达的地方和如何产生数据。如果你感觉一个方法 / 函数负责太多功能了（生成数据并且写入一个数据库？处理 HTTP 请求并且处理业务级别的逻辑），那么你可能就需要 DI 这个工具了。
在不同环境下重用代码。我们的代码所处的第一个「新」环境就是在内部进行测试。但是随后，如果其他人想要用你的代码尝试点新东西，他们只要注入他们自己的依赖就可以了。
*/
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
