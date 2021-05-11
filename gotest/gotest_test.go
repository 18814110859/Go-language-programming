package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, _ := Division(6, 2); i != 3 {
		t.Error("测试没通过！")
	} else {
		t.Log("测试通过！")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Division(6, 2)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	// 调用该函数停止压力测试的时间计数
	b.StopTimer()

	// 重新开始时间
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Division(6, 2)
	}
}
