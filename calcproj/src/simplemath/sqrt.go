package simplemath

import "math"
// 计算 i 的 平方根
func Sqrt (i int) int {
	f := float64(i)
	v := math.Sqrt(f)
	return int(v)
}
