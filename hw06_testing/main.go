package main

import (
	"fmt"

	"github.com/n0more/otus_hw/hw06_testing/hw"
)

const (
	Year int = iota
	Size int = iota
	Rate int = iota
)

func main() {
	outputHw02 := hw.Runhw02("data.json")
	fmt.Println(outputHw02)
	outputHw03 := hw.Runhw03(10)
	fmt.Println(outputHw03)

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

	outputHw041 := hw.Runhw04(bookFirst, bookSecond, Year)
	fmt.Println(outputHw041)
	outputHw042 := hw.Runhw04(bookFirst, bookSecond, Rate)
	fmt.Println(outputHw042)
	outputHw043 := hw.Runhw04(bookFirst, bookSecond, Size)
	fmt.Println(outputHw043)

	circle := hw.Circle{Radius: 5}
	rectangle := hw.Rectangle{Width: 10, Height: 5}
	triangle := hw.Triangle{Base: 8, Height: 6}
	rhombus := "test"

	objects := []any{circle, rectangle, triangle, rhombus}

	for _, obj := range objects {
		result, err := hw.Runhw05(obj)
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
		} else {
			fmt.Println(result)
		}
	}
}
