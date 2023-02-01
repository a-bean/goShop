package monkey

import (
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
	"testing"
)

// mock 一个函数, mock 范围更广， 而且不需要事先生成代码， 大家可以结合自己的需求
func TestCompute(t *testing.T) {
	//动态的补丁技术
	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 2, nil
	})
	defer patches.Reset()

	sum, err := Compute(1, 2)
	if err != nil {
		t.Error(err)
	}
	if sum != 3 {
		t.Errorf("sum is %d, want 3", sum)
	}
}

func TestCompute2(t *testing.T) {
	//动态的补丁技术
	var c *Computer
	patches := gomonkey.ApplyMethod(reflect.TypeOf(c), "NetworkCompute", func(_ *Computer, a, b int) (int, error) {
		return 2, nil
	})
	defer patches.Reset()

	c = &Computer{}
	sum, err := c.Compute(1, 2)
	if err != nil {
		t.Error(err)
	}
	if sum != 3 {
		t.Errorf("sum is %d, want 3", sum)
	}
}

var num = 10

func TestGlobalVar(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&num, 12)
	defer patches.Reset()

	if num != 10 {
		t.Errorf("expected %v, got: %v", 10, num)
	}
}
