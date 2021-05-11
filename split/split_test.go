package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("测试之前的setup")
	return func(t *testing.T) {
		t.Log("测试之后的teardown")
	}
}

// 自测试的 setup 与 teardown
func setupTestSub(t *testing.T) func(t *testing.T) {
	t.Log("子测试之前的setup")
	return func(t *testing.T) {
		t.Log("子测试之前的teardown")
	}
}

func TestSplit(t *testing.T) {

	result := Split("a:b:c", ":")
	t1 := []string{"a", "b", "c"}
	if !reflect.DeepEqual(result, t1) {
		t.Errorf("expect:%v, result:%v", t1, result)
	}
}

func TestMoreSplit(t *testing.T) {
	result := Split("abcd", "bc")
	t1 := []string{"a", "d"}
	if !reflect.DeepEqual(result, t1) {
		t.Errorf("expect:%v, result:%v", t1, result)
	}
}

func TestSubSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}

	teardownTestCase := setupTestCase(t) // setup 操作
	defer teardownTestCase(t)            // teardown 操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			teardownTestSub := setupTestSub(t) // setup 操作
			defer teardownTestSub(t)           // teardown 操作

			ret := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("expect:%#v, result:%#v", tc.want, ret)
			}
		})
	}

}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.SetParallelism(16) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("枯藤老树昏鸦", "老")
		}
	})
}
