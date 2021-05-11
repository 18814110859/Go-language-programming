package split

import (
	"strings"
)

func Split(s, sep string) (result []string) {
	// 提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加.
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}

	result = append(result, s)
	return
}
