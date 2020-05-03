package main

import "fmt"

type Strategy interface {
	SumN(n int)
}

type LoopSumStrategy struct{}

func (l *LoopSumStrategy) SumN(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("loop sum:", sum)
}

type FormulaSumStrategy struct {
}

func (f *FormulaSumStrategy) SumN(n int) {
	fmt.Println("formula sum:", n*(n+1)/2)
}

type Solver struct {
	n int
	s Strategy
}

func (s *Solver) Do() {
	s.s.SumN(s.n)
}

func main() {
	s := Solver{n: 100, s: &LoopSumStrategy{}}
	s.Do()

	s.s = &FormulaSumStrategy{}
	s.Do()
}
