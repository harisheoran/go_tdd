package main

import "math"

type Shape interface {
	Area() float64
}

// like class in Java
type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// methods
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func Perimeter(rectable Rectangle) float64 {
	return 2 * (rectable.Length + rectable.Width)
}
