package hw

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

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float32 {
	return b.rate
}

const (
	Year int = iota
	Size int = iota
	Rate int = iota
)

type Filter struct {
	field int
}

func NewCompare(field int) *Filter {
	return &Filter{field: field}
}

func (f *Filter) Compare(b1, b2 Book) bool {
	switch f.field {
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

func Runhw04(bookFirst Book, bookSecond Book, field int) string {
	Compare := NewCompare(field)

	return fmt.Sprint(Compare.Compare(bookFirst, bookSecond))
}
