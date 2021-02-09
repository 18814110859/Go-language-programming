package simplemath

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	ret := Multiply(2, 2)
	if ret != 4 {
		t.Errorf("Multiply(2, 2) failed. Got %v, expected 4.", ret)
	}
}