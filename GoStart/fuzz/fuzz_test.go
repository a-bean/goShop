package fuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

/*
模糊测试，单元测试有局限性， 每个测试输入必须由开发者添加到单元测试用例， 更进一步
还是一个人工的check过程
fuzzing优点 就是可以基于开发者代码中指定的测试输入作为基础数据，进一步自动生成随机测试数据，用来发现指定测试输入没有覆盖到的边界情况
*/

func Reverse(s string) (string, error) {
	//没有考虑到非法的unicode编码
	//q 代表该值对应的单引号括起来的go语法字符字面值
	if !utf8.ValidString(s) {
		return "", fmt.Errorf("invalid utf8: %q", s)
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

//TDD
/*
当涉及业务相关的单元测试, 此时

用例复杂, 输入可能是多层嵌套的struct, 某一层的某个变量的值影响输出
用例数量多, 某个函数可能有5个以上的用例
由于table-driven的表达能力有限:

如何复用输入结构体? (当前情况开发会复制粘贴过去改)
name过于简单不被重视, 导致单元测试失败时难以快速阅读

最终带来的问题是:

单个用例构造复杂, 可能是十几行甚至是几十行;
用例和用例之间没有复用, 基本基于复制粘贴;
难以区分用例之间的差异
用例过多导致难以维护, 不能明确知道每个用例的目的, 用例和用例之间的差别;
单个测试集过大, 可能有几百行测试代码;
BDD
goconvey ginkgo
*/

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"hello", "olleh"}, //基本
		{"a", "a"},         //边界
		{" ", " "},         //特殊
	}
	for _, c := range testcases {
		rev, _ := Reverse(c.in)
		assert.Equal(t, c.want, rev)
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello", "a", " ", "!#j"}
	for _, c := range testcases {
		f.Add(c) //提供种子语料库
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}
		//没有输入的，也有无法预期的输出， 模糊测试的缺点非常明显，就是几乎无法通过equal去判断， 只能通过比较字符串的长度来判断
		//如果要基于种子语料库生成随机测试数据用于模糊测试，需要给go test命令加上 -fuzz=Fuzz
		assert.Equal(t, len(orig), len(rev))
		double, err := Reverse(rev)
		if err != nil {
			return
		}
		assert.Equal(t, orig, double) //技巧

		//另一种技巧，这个技巧说实话没有什么说服力
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("invalid utf8: %q", rev)
		}
	})
}

//不要把单元和模糊测试做对比，他们是互补的，某些核心函数我们可以同时将单元测试和模糊加上保证代码的正确性

//testdata目录go的一种惯例， 如果我们的测试用例有外部输入数据，我们就将数据放到testdata，这个目录会被go test命令扫描到， 这样我们就不用关心路径问题

//protoc 执行过程 会从标注输入 读取到你的参数 回去查询 protoc-gen-{NAME} go_out 会去找 protoc-gen-go.exe
