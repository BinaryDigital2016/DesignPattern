package main

import "fmt"

type ProductA interface {
	Typing()
	PlayVideo(name string)
}

type ProductA1 struct {
}

func (p *ProductA1) Typing() {
	fmt.Println("ProductA1 is typing...")
}

func (p *ProductA1) PlayVideo(name string) {
	fmt.Printf("ProductA1 is playing video:%s\n", name)
}

type ProductA2 struct {
}

func (p *ProductA2) Typing() {
	fmt.Println("ProductA2 is typing...")
}

func (p *ProductA2) PlayVideo(name string) {
	fmt.Printf("ProductA2 is playing video:%s\n", name)
}

type ProductB interface {
	Eating()
	WatchVideo(name string)
}

type ProductB1 struct {
}

func (p *ProductB1) Eating() {
	fmt.Println("ProductB1 is eating...")
}

func (p *ProductB1) WatchVideo(name string) {
	fmt.Printf("ProductB1 is watching video:%s\n", name)
}

type ProductB2 struct {
}

func (p *ProductB2) Eating() {
	fmt.Println("ProductB2 is eating...")
}

func (p *ProductB2) WatchVideo(name string) {
	fmt.Printf("ProductB2 is watching video:%s\n", name)
}

type IFactory interface {
	CreateA() ProductA
	CreateB() ProductB
}

type Factory1 struct {
}

func (f *Factory1) CreateA() ProductA {
	return &ProductA1{}
}

func (f *Factory1) CreateB() ProductB {
	return &ProductB1{}
}

type Factory2 struct {
}

func (f *Factory2) CreateA() ProductA {
	return &ProductA2{}
}

func (f *Factory2) CreateB() ProductB {
	return &ProductB2{}
}

func main() {
	var f IFactory
	f = &Factory1{}
	a := f.CreateA()
	a.Typing()
	a.PlayVideo("西游记")

	b := f.CreateB()
	b.Eating()
	b.WatchVideo("西游记")

	f = &Factory2{}
	a = f.CreateA()
	a.Typing()
	a.PlayVideo("西游记")

	b = f.CreateB()
	b.Eating()
	b.WatchVideo("西游记")
}
