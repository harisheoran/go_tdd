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
	t.Run("Area of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		expected := rectangle.Area()
		result := 200.0

		if expected != result {
			t.Errorf("expected %.2f and result %.2f", expected, result)
		}
	})

	t.Run("Area of Circle", func(t *testing.T) {
		circle := Circle{4.0}
		expected := circle.Area()
		result := 50.26548245743668986

		if expected != result {
			t.Errorf("expected %.17f and result %.17f", expected, result)
		}

	})

}
