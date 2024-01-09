package main

import "math"

// like class in Java
type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}

func Perimeter(rectable Rectangle) float64 {
	return 2 * (rectable.Length + rectable.Width)
}
