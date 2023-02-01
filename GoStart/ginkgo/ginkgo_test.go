package ginkgo

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

// 并不一定需要每个测试用例都这么写， 对于核心的函数或者核心的业务逻辑我们建议设计好的测试用例
var _ = Describe("Books", func() {
	var (
		longBook  string
		shortBook string

		pathches *gomonkey.Patches
		ctl      *gomock.Controller
	)
	BeforeEach(func() {
		longBook = "long"
		shortBook = "short"

		ctl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		longBook = ""
		shortBook = ""

		ctl.Finish()
		pathches.Reset()
	})

	Describe("Add Books", func() {
		It("should be able to add a book", func() {
			//调用AddBook方法，并传入参数，期望返回的结果为true
			assert.Equal(GinkgoT(), "long", longBook)
		})
		It("should not be able to add a book", func() {
			//调用AddBook方法，并传入参数，期望返回的结果为true
			assert.Equal(GinkgoT(), "short", shortBook)
		})
	})

	Describe("Delete Books", func() {

	})
})

/*
1. proto文件可以用作http和rpc服务的生成标注写法
	我写了一个gin的服务，我还要手动去维护api文档，手动去yapi上维护 后期维护和迭代很简单， 改了任何代码你都可以直接生成api
	可以直接将proto生成swagger文件，然后一键导入到yapi上，这样就可以直接在yapi上查看api文档了
2. 在kratos中对proto的依赖更加重， 可以用来定义一些错误码， 并生成go源码直接使用
3. kratos甚至将配置文件都给你映射成proto文件
业内很多框架都开始逐步接受将proto文件作为核心的标准去写一系列插件去自动生成代码
proto validate

go-zero更溜，goctl，保姆式的框架 api文件 go-zero和kratos的一套设计理念
*/
