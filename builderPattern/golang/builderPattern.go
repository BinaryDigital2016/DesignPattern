package main

import "fmt"

type BMW struct {
}

func (b *BMW) Race() {
	fmt.Println("race with BMW")
}

type BYD struct {
}

func (b *BYD) Taxi() {
	fmt.Println("taxi with BYD")
}

type Builder interface {
	BuildEngine()
	BuildWheel()
	BuildNavigation()
	GetCar() interface{}
}

type BMWBuilder struct {
}

func (b *BMWBuilder) BuildEngine() {
	fmt.Println("BMW build engine")
}

func (b *BMWBuilder) BuildWheel() {
	fmt.Println("BMW build wheel")
}

func (b *BMWBuilder) BuildNavigation() {
	fmt.Println("BMW build navigation")
}

func (b *BMWBuilder) GetCar() interface{} {
	return BMW{}
}

type BYDBuilder struct {
}

func (b *BYDBuilder) BuildEngine() {
	fmt.Println("BYD build engine")
}

func (b *BYDBuilder) BuildWheel() {
	fmt.Println("BYD build wheel")
}

func (b *BYDBuilder) BuildNavigation() {
	//fmt.Println("BYD build navigation")
}

func (b *BYDBuilder) GetCar() interface{} {
	return BYD{}
}

// direct the build process
type Director struct {
	b Builder
}

func (d *Director) Build() {
	d.b.BuildEngine()
	d.b.BuildWheel()
	d.b.BuildNavigation()
}

func main() {
	bydBuilder := BYDBuilder{}
	d := Director{b: &bydBuilder}
	d.Build()
	byd := bydBuilder.GetCar().(BYD)
	byd.Taxi()

	bmwBuilder := BMWBuilder{}
	d = Director{b: &bmwBuilder}
	d.Build()
	bmw := bmwBuilder.GetCar().(BMW)
	bmw.Race()
}
