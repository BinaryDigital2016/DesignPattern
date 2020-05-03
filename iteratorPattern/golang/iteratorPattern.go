package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type SliceIterator struct {
	cur int
	sc  *SliceContainer
}

func (s *SliceIterator) HasNext() bool {
	if len(s.sc.elem) > s.cur {
		return true
	}

	return false
}

func (s *SliceIterator) Next() interface{} {
	if s.HasNext() {
		defer func() {
			s.cur++
		}()
		return s.sc.elem[s.cur]
	}
	return nil
}

type Container interface {
	GetIterator() Iterator
}

type SliceContainer struct {
	elem []interface{}
}

func (s *SliceContainer) GetIterator() Iterator {
	return &SliceIterator{
		cur: 0,
		sc:  s,
	}
}

func main() {
	sc := SliceContainer{
		elem: []interface{}{1, 2, 3, 4},
	}
	iter := sc.GetIterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
