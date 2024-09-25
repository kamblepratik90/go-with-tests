package main

import "testing"

type Rectangle struct {
	Length float64
	Width  float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Length + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Length * rect.Width
}

func TestPerimeter(t *testing.T) {
	var rect Rectangle = Rectangle{10.3, 19.3}
	perimeter := Perimeter(rect)
	expectedPeri := 59.2

	if perimeter != expectedPeri {
		t.Errorf("perimeter is %.2f and expectedPeri is %.2f", perimeter, expectedPeri)
	}
}

func TestArea(t *testing.T) {
	var rect Rectangle = Rectangle{10.0, 10.0}
	area := Area(rect)
	expectedArea := 100.0

	if area != expectedArea {
		t.Errorf("area is %.2f and expectedArea is %.2f", area, expectedArea)
	}
}
