package gotest

import "errors"

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为0！")
	}
	return a / b, nil
}
