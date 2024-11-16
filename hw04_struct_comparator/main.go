package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b Book) GetID() int {
	return b.id
}

func (b Book) GetTitle() string {
	return b.title
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b Book) GetYear() int {
	return b.year
}

func (b Book) GetSize() int {
	return b.size
}

func (b Book) GetRate() float32 {
	return b.rate
}

const (
	Year int = iota
	Size int = iota
	Rate int = iota
)

func Compare(b1, b2 Book, filter int) bool {
	switch filter {
	case Year:
		return b1.year > b2.year
	case Size:
		return b1.size > b2.size
	case Rate:
		return b1.rate > b2.rate
	default:
		return false
	}
}

func main() {
	bookFirst := Book{}
	bookFirst.SetID(1)
	bookFirst.SetAuthor("Book1 Author")
	bookFirst.SetTitle("Хорошая Книга")
	bookFirst.SetRate(1.4)
	bookFirst.SetSize(150)
	bookFirst.SetYear(1956)

	bookSecond := Book{}
	bookSecond.SetID(2)
	bookSecond.SetAuthor("Author unknown")
	bookSecond.SetTitle("Умная Книга")
	bookSecond.SetRate(1.3)
	bookSecond.SetSize(100)
	bookSecond.SetYear(2956)

	fmt.Println(Compare(bookFirst, bookSecond, Year))
	fmt.Println(Compare(bookFirst, bookSecond, Size))
	fmt.Println(Compare(bookFirst, bookSecond, Rate))
}
