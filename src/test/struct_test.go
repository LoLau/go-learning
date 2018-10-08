package test

import (
	"testing"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Hight float64
}

// 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Hight / 2
}

func TestArea(t *testing.T) {
	//assertCorrectMessage := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f, want %.2f", got, want)
	//	}
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{2, 3}
	//	want := 6.0
	//	assertCorrectMessage(t, rectangle, want)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	want := 314.1592653589793
	//	assertCorrectMessage(t, circle, want)
	//})

	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 361.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}

// 函数
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
