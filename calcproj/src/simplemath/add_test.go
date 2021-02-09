package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 2 {
		t.Errorf("Add(1, 2) failed. got %d, expected 3.", r)
	}
}