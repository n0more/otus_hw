package hw

import (
	"errors"
	"fmt"
	"strings"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	const pi = 3.141592653589793115997963468544185161590576171875
	return pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}

func Runhw05(obj any) (string, error) {
	output := []string{}
	area, err := calculateArea(obj)
	if err != nil {
		return "error", err
	}
	switch shape := obj.(type) {
	case Circle:
		output = append(output, fmt.Sprintf("Круг: радиус %v\n", shape.Radius))
		output = append(output, fmt.Sprintf("Площадь круга: %.14f\n", area))
	case Rectangle:
		output = append(output, fmt.Sprintf("Прямоугольник: ширина %v, высота %v\n", shape.Width, shape.Height))
		output = append(output, fmt.Sprintf("Площадь прямоугольника: %v\n", area))
	case Triangle:
		output = append(output, fmt.Sprintf("Треугольник: основание %v, высота %v\n", shape.Base, shape.Height))
		output = append(output, fmt.Sprintf("Площадь треугольника: %v\n", area))
	default:
		output = append(output, fmt.Sprintln("Неизвестная фигура"))
	}
	return strings.Join(output, ""), nil
}
