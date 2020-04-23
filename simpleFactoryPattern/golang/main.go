package main

import "fmt"

type ProductType string

const (
	ProductTypeA ProductType = "A"
	ProductTypeB ProductType = "B"
)

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

func Create(t ProductType) IProduct {
	switch t {
	case ProductTypeA:
		return &ProductA{}
	case ProductTypeB:
		return &ProductB{}
	default:
		return &ProductA{}
	}
}

func main() {
	a := Create(ProductTypeA)
	a.Typing()
	a.PlayVideo("西游记")

	b := Create(ProductTypeB)
	b.Typing()
	b.PlayVideo("西游记")
}
