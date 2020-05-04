package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Paper struct {
	Title string
	Size  string
}

func (p *Paper) Clone() Prototype {
	cp := Paper{Title: p.Title, Size: p.Size}
	return &cp
}

func main() {
	p := Paper{Title: "本科毕业论文", Size: "A4"}
	fmt.Printf("p:%+v\n", p)

	p1 := p.Clone().(*Paper)
	fmt.Printf("before modify: p1:%+v\n", p1)
	p1.Title = "研究生毕业论文"
	fmt.Printf("after modify: p1:%+v\n", p1)
	fmt.Printf("after modify: p:%+v\n", p)
}
