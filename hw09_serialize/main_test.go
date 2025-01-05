package main

import (
	"testing"

	"github.com/n0more/otus_hw/hw09_serialize/format"
	"github.com/stretchr/testify/assert"
)

func Test_mJson(t *testing.T) {
	book := format.Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2021,
		Size:   350,
		Rate:   4.5,
		Sample: []byte("SomeBytes"),
	}

	jsonData, err := book.MarshalJSON()
	assert.Nil(t, err)
	result := `{"id":1,"title":"Go Programming","author":"John Doe","year":2021,"size":350,` +
		`"rate":4.5,"sample":"U29tZUJ5dGVz"}`
	assert.Equal(t, result, string(jsonData))
}

func Test_mXml(t *testing.T) {
	book := format.Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2021,
		Size:   350,
		Rate:   4.5,
		Sample: []byte("SomeBytes"),
	}

	xmlData, err := book.WriteXML()
	assert.Nil(t, err)
	result := `<Book><id>1</id><title>Go Programming</title><author>John Doe</author><year>2021</year>` +
		`<size>350</size><rate>4.5</rate><sample>SomeBytes</sample></Book>`
	assert.Equal(t, result, string(xmlData))
}
