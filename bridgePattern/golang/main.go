package main

import "fmt"

type Worker interface {
	Leaving()
}

type GoodWorker struct{}

func (g *GoodWorker) Leaving() {
	fmt.Println("好员工离职")
}

type NormalWorker struct{}

func (n *NormalWorker) Leaving() {
	fmt.Println("普通员工离职")
}

type Company interface {
	Running()
}

type BigCompany struct {
	Worker
}

func (b *BigCompany) Running() {
	fmt.Println("每个员工都是螺丝钉")
	b.Worker.Leaving()
	fmt.Println("正常运行")
}

type SmallCompany struct {
	Worker
}

func (s *SmallCompany) Running() {
	fmt.Println("每个员工都是骨干")
	s.Worker.Leaving()
	fmt.Println("业务受阻")
}

func main() {
	gw := &GoodWorker{}
	nw := &NormalWorker{}
	bc := &BigCompany{Worker: gw}
	bc.Running()
	bc2 := &BigCompany{Worker: nw}
	bc2.Running()
	sc := &SmallCompany{Worker: gw}
	sc.Running()
	sc2 := &SmallCompany{Worker: nw}
	sc2.Running()
}
