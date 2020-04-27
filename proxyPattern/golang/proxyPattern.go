package main

import "fmt"

type Doer interface {
	DoSomething()
}

type ConcreteDoer struct {
	Name string
}

func (c *ConcreteDoer) DoSomething() {
	fmt.Printf("%s do something", c.Name)
}

type Proxy struct {
	Doer
}

func (p *Proxy) DoSomething() {
	fmt.Println("我是一个代理")
	p.Doer.DoSomething()
}

func main() {
	d := ConcreteDoer{Name: "Frank"}
	p := Proxy{Doer: &d}
	p.DoSomething()
}
