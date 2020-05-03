package main

import "fmt"

type Element interface {
	Accept(v Visitor)
	Name() string
}

type Company struct {
	CName string
}

// 修改v的SendNewspaper方法即可修改Company的表现
func (c *Company) Accept(v Visitor) {
	v.SendElement(c)
}

func (c *Company) Name() string {
	return c.CName
}

type School struct {
	SName string
}

func (s *School) Accept(v Visitor) {
	v.SendElement(s)
}

func (s *School) Name() string {
	return s.SName
}

type Visitor interface {
	SendElement(e Element)
}

type Milkman struct {
}

func (m *Milkman) SendElement(e Element) {
	fmt.Printf("milkman send milk to %s\n", e.Name())
}

type DeliveryBoy struct {
}

func (d *DeliveryBoy) SendElement(e Element) {
	fmt.Printf("deliverboy send newspaper to %s\n", e.Name())
}

func main() {
	s := School{SName: "School A"}
	s.Accept(&Milkman{})

	c := Company{CName: "Company B"}
	c.Accept(&DeliveryBoy{})
}
