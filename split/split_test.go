package split

import (
	"reflect"
	"testing"
)

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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
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
