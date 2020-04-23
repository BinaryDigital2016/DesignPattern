package main

import "fmt"

type IProduct interface {
	Typing()
	PlayVideo(name string)
}

type ProductA struct {
}

func (p *ProductA) Typing() {
	fmt.Println("ProductA is typing...")
}

func (p *ProductA) PlayVideo(name string) {
	fmt.Printf("ProductA is playing video:%s\n", name)
}

type ProductB struct {
}

func (p *ProductB) Typing() {
	fmt.Println("ProductB is typing...")
}

func (p *ProductB) PlayVideo(name string) {
	fmt.Printf("ProductB is playing video:%s\n", name)
}

type IFactory interface {
	Create() IProduct
}

type FactoryA struct {
}

func (f *FactoryA) Create() IProduct {
	return &ProductA{}
}

type FactoryB struct {
}

func (f *FactoryB) Create() IProduct {
	return &ProductB{}
}

func main() {
	var f IFactory
	f = &FactoryA{}
	a := f.Create()
	a.Typing()
	a.PlayVideo("西游记")

	f = &FactoryB{}
	b := f.Create()
	b.Typing()
	b.PlayVideo("西游记")
}
