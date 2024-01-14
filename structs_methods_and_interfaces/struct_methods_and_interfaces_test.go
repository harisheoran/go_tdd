package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	expected := Perimeter(rectangle)
	result := 60.0

	if expected != result {
		t.Errorf("expected %.2f and result %.2f", expected, result)
	}
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		name   string
		shape  Shape
		result float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 20.0}, result: 200.0},
		{name: "Circle", shape: Circle{4.0}, result: 50.26548245743668986},
		{name: "Triangle", shape: Triangle{12, 6}, result: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.shape.Area()
			if expected != tt.result {
				t.Errorf("got %g want %g", expected, tt.result)
			}
		})
	}
}
