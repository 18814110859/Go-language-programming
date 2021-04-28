package main

import "errors"

type IntSet struct {
	data map[int]bool
}


func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int)  {
	set.data[x] = true
}


func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	_, ok := set.data[x]
	return ok
}

type UndoableIntSet struct {
	IntSet
	functions []func()
}


func NewUndoableIntSet () UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) {
	if !set.IntSet.Contains(x) {
		set.IntSet.data[x] = true
		set.functions = append(set.functions, func() { set.IntSet.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}


func (set *UndoableIntSet) Delete(x int) {
	if set.IntSet.Contains(x) {
		delete(set.IntSet.data, x)
		set.functions = append(set.functions, func() { set.IntSet.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("no function to undo")
	}

	lastIndex := len(set.functions) - 1
	if function := set.functions[lastIndex]; function != nil {
		function()
		set.functions[lastIndex] = nil
	}

	set.functions = set.functions[:lastIndex]

	return nil
}







