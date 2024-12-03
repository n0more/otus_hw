package main

import (
	"testing"

	"github.com/n0more/otus_hw/hw06_testing/hw"
	"github.com/stretchr/testify/assert"
)

func Test_hw02(t *testing.T) {
	// test with default data.json file
	output := hw.Runhw02("data.json")

	expected := "User ID: 10; Age: 25; Name: Rob; Department ID: 3; " +
		"\nUser ID: 11; Age: 30; Name: George; Department ID: 2; "

	assert.Equal(t, expected, output)
}

func Test_hw03(t *testing.T) {
	// test with size 5

	output := hw.Runhw03(5)

	expected := "Размер доски: 5 на 5\n#_#_#\n_#_#_\n#_#_#\n_#_#_\n#_#_#\n"

	assert.Equal(t, expected, output)
}

func Test_hw04(t *testing.T) {
	// test with size 5

	bookFirst := hw.Book{}
	bookFirst.SetID(1)
	bookFirst.SetAuthor("Book1 Author")
	bookFirst.SetTitle("Хорошая Книга")
	bookFirst.SetRate(1.4)
	bookFirst.SetSize(150)
	bookFirst.SetYear(1956)

	bookSecond := hw.Book{}
	bookSecond.SetID(2)
	bookSecond.SetAuthor("Author unknown")
	bookSecond.SetTitle("Умная Книга")
	bookSecond.SetRate(1.7)
	bookSecond.SetSize(100)
	bookSecond.SetYear(956)

	testCases := []struct {
		name   string
		fBook  hw.Book
		sBook  hw.Book
		filter int
		result string
	}{
		{
			name:   "Year Compare",
			fBook:  bookFirst,
			sBook:  bookSecond,
			filter: hw.Year,
			result: "true",
		},
		{
			name:   "Rate Compare",
			fBook:  bookFirst,
			sBook:  bookSecond,
			filter: hw.Rate,
			result: "false",
		},
		{
			name:   "Size Compare",
			fBook:  bookFirst,
			sBook:  bookSecond,
			filter: hw.Size,
			result: "true",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := hw.Runhw04(tC.fBook, tC.sBook, tC.filter)
			assert.Equal(t, tC.result, got)
		})
	}
}

func Test_hw05(t *testing.T) {
	testCases := []struct {
		name   string
		shape  hw.Shape
		result string
	}{
		{
			name:   "Circle Area",
			shape:  hw.Circle{Radius: 5},
			result: "Круг: радиус 5\nПлощадь круга: 78.53981633974483\n",
		},
		{
			name:   "Rectangle Area",
			shape:  hw.Rectangle{Width: 10, Height: 5},
			result: "Прямоугольник: ширина 10, высота 5\nПлощадь прямоугольника: 50\n",
		},
		{
			name:   "Triangle Area",
			shape:  hw.Triangle{Base: 8, Height: 6},
			result: "Треугольник: основание 8, высота 6\nПлощадь треугольника: 24\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := hw.Runhw05(tC.shape)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tC.result, got)
		})
	}
}
