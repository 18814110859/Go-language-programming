package main

import (
	"fmt"
	"math"
)

func main() {
	traversalString()

}

func traversalString() {

	s := "hello world go.你好 go"
	for i := 0; i < len(s); i++ {
		// byte 或者 理解为 char
		fmt.Printf("%v{%c}\n", s[i], s[i])
	}
	fmt.Println()
	for _, v := range s {
		// rune
		fmt.Printf("%v{%c}\n", v, v)
	}

	fmt.Println("修改字符串代码～～～")
	s1 := "hello go"
	bytes := []byte(s1)
	bytes[0] = 'H'
	for k, v := range bytes {
		fmt.Printf("k:%d v:%v{%c}{%p}\n", k, v, v, &v)
	}
	// 输出结果
	// k:0 v:72{H}{0xc00012c038}
	// k:1 v:101{e}{0xc00012c038}
	// k:2 v:108{l}{0xc00012c038}
	// k:3 v:108{l}{0xc00012c038}
	// k:4 v:111{o}{0xc00012c038}
	// k:5 v:32{ }{0xc00012c038}
	// k:6 v:103{g}{0xc00012c038}
	// k:7 v:111{o}{0xc00012c038}

	fmt.Printf("s1: %s{%p} bytes: %s{%p}\n", s1, &s1, string(bytes), &bytes)
	// s1: hello go{0xc0001041e0} bytes: Hello go{0xc00011c020}

	s2 := "今天大涨"
	runes := []rune(s2)
	runes[3] = '跌'
	fmt.Printf("s2: %s{%p} runes: %s{%p}\n", s2, &s2, string(runes), &runes)
	// s2: 中国{0xc000104210} runes: 美国{0xc00011c040}

	fmt.Printf("runes[0]: %s{%c}{%p}\n", string(runes[0]), runes[0], &runes[0])
	fmt.Printf("runes[1]: %s{%c}{%p}\n", string(runes[1]), runes[1], &runes[1])
	fmt.Printf("runes[2]: %s{%c}{%p}\n", string(runes[2]), runes[2], &runes[2])
	fmt.Printf("runes[3]: %s{%c}{%p}\n", string(runes[3]), runes[3], &runes[3])
	// runes[0]: 今{今}{0xc0000b4050}
	// runes[1]: 天{天}{0xc0000b4054}
	// runes[2]: 大{大}{0xc0000b4058}
	// runes[3]: 跌{跌}{0xc0000b405c}

	fmt.Println(strQrtDemo())

}

/**
 * 1.1.11. 类型转换
 * Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
 * 强制类型转换的基本语法如下：
 * T(表达式)
 * 其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
 * 比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
 */
func strQrtDemo() int {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	x := a*a + b*b
	c = int(math.Sqrt(float64(x)))
	return c
}
