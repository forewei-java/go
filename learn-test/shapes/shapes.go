package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area 类方法计算面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 周长
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

type Circle struct {
	Radius float64
}

// Area 面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

type Shape interface {
	Area() float64
}
