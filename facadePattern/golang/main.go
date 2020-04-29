package main

import "fmt"

type SubSystemOne struct {
}

func (s *SubSystemOne) MethodOne() {
	fmt.Println("sub system method one")
}

type SubSystemTwo struct {
}

func (s *SubSystemTwo) MethodTwo() {
	fmt.Println("sub system method two")
}

type SubSystemThree struct {
}

func (s *SubSystemThree) MethodThree() {
	fmt.Println("sub system method three")
}

type SubSystemFour struct {
}

func (s *SubSystemFour) MethodFour() {
	fmt.Println("sub system method four")
}

type IFacade interface {
	MethodA()
	MethodB()
}

type ConcreteFacade struct {
	SubOne   *SubSystemOne
	SubTwo   *SubSystemTwo
	SubThree *SubSystemThree
	SubFour  *SubSystemFour
}

func (c *ConcreteFacade) MethodA() {
	c.SubOne.MethodOne()
	c.SubTwo.MethodTwo()
}

func (c *ConcreteFacade) MethodB() {
	c.SubThree.MethodThree()
	c.SubFour.MethodFour()
}

func main() {
	f := ConcreteFacade{
		SubOne:   &SubSystemOne{},
		SubTwo:   &SubSystemTwo{},
		SubThree: &SubSystemThree{},
		SubFour:  &SubSystemFour{},
	}

	f.MethodA()
	f.MethodB()
}
