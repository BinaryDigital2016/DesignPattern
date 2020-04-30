package main

import "fmt"

const (
	MsgPattern = "%s 批准了 %s 加薪 %d 元的请求。\n"
)

type Manager interface {
	Approve(name string, num int)
}

type CommonManager struct {
	Name     string
	Limit    int
	Superior Manager
}

func (c *CommonManager) Approve(name string, num int) {
	if num > c.Limit {
		if c.Superior == nil {
			panic("权限不够")
		}
		c.Superior.Approve(name, num)
		return
	}

	fmt.Printf(MsgPattern, c.Name, name, num)
}

type Majordomo struct {
	Name     string
	Limit    int
	Superior Manager
}

func (m *Majordomo) Approve(name string, num int) {
	if num > m.Limit {
		if m.Superior == nil {
			panic("权限不够")
		}
		m.Superior.Approve(name, num)
		return
	}

	fmt.Printf(MsgPattern, m.Name, name, num)
}

type GeneralManager struct {
	Name     string
	Limit    int
	Superior Manager
}

func (g *GeneralManager) Approve(name string, num int) {
	fmt.Printf(MsgPattern, g.Name, name, num)
}

func main() {
	g := GeneralManager{
		Name: "Frank",
	}

	m := Majordomo{
		Name:     "Ronnie",
		Limit:    5000,
		Superior: &g,
	}

	c := CommonManager{
		Name:     "Allen",
		Limit:    1000,
		Superior: &m,
	}

	c.Approve("aaa", 500)
	c.Approve("bbb", 3000)
	c.Approve("ccc", 10000)
}
