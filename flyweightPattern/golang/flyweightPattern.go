package main

import (
	"fmt"
	"math/rand"
)

type Color string

const (
	Green Color = "green"
	Red   Color = "red"
	Black Color = "black"
)

//type Shape interface {
//	SetX(x int)
//	SetY(y int)
//	SetRadius(r int)
//	SetColor()
//	Draw()
//}

type Circle struct {
	x      int
	y      int
	radius int
	color  Color
}

func (c *Circle) Circle(color Color) {
	c.color = color
}

func (c *Circle) SetX(x int) {
	c.x = x
}

func (c *Circle) SetY(Y int) {
	c.y = Y
}

func (c *Circle) SetRadius(r int) {
	c.radius = r
}

func (c *Circle) Draw() {
	fmt.Printf("Circle: Draw() [Color:%v, x:%d, y:%d, radius:%d]\n", c.color, c.x, c.y, c.radius)
}

type ShapeFactory struct {
	circles map[Color]*Circle
}

// 暂不考虑并发
func (s *ShapeFactory) GetCircle(c Color) *Circle {
	if s.circles == nil {
		s.circles = make(map[Color]*Circle)
	}

	circle, ok := s.circles[c]
	if !ok {
		circle = new(Circle)
		circle.Circle(c)
		s.circles[c] = circle
	}

	return circle
}

type FlyweightPatternRandom struct {
	colors []Color
}

func (f *FlyweightPatternRandom) GetRandomColor() Color {
	if f.colors == nil {
		f.colors = []Color{Green, Red, Black}
	}

	num := rand.Intn(len(f.colors) - 1)
	return f.colors[num]
}

func (f *FlyweightPatternRandom) Random() int {
	num := rand.Intn(10) * 100
	return num
}

func main() {
	f := FlyweightPatternRandom{}
	shapeFactory := new(ShapeFactory)
	for i := 0; i < 10; i++ {
		c := shapeFactory.GetCircle(f.GetRandomColor())
		c.SetX(f.Random())
		c.SetY(f.Random())
		c.SetRadius(10)
		fmt.Println("------------------")
		c.Draw()
	}
}
