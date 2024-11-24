package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	const pi = 3.141592653589793115997963468544185161590576171875
	return pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	triangle := Triangle{base: 8, height: 6}
	rhombus := "test"

	objects := []any{circle, rectangle, triangle, rhombus}

	for _, obj := range objects {
		area, err := calculateArea(obj)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			switch obj.(type) {
			case Circle:
				fmt.Printf("Круг: радиус %v\n", circle.radius)
				fmt.Printf("Площадь круга: %.14f\n", area)
			case Rectangle:
				fmt.Printf("Прямоугольник: ширина %v, высота %v\n", rectangle.width, rectangle.height)
				fmt.Printf("Площадь прямоугольника: %v\n", area)
			case Triangle:
				fmt.Printf("Треугольник: основание %v, высота %v\n", triangle.base, triangle.height)
				fmt.Printf("Площадь треугольника: %v\n", area)
			default:
				fmt.Println("Неизвестная фигура")
			}
		}
	}
}
