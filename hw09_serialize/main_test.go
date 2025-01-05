package main

import (
	"testing"

	"github.com/n0more/otus_hw/hw09_serialize/format"
	"github.com/stretchr/testify/assert"
)

var book = format.Book{
	ID:     1,
	Title:  "Go Programming",
	Author: "John Doe",
	Year:   2021,
	Size:   350,
	Rate:   4.5,
	Sample: []byte("SomeBytes"),
}

func Test_Json(t *testing.T) {
	jsonData, err := book.MarshalJSON()
	assert.Nil(t, err)
	result := `{"id":1,"title":"Go Programming","author":"John Doe","year":2021,"size":350,` +
		`"rate":4.5,"sample":"U29tZUJ5dGVz"}`
	assert.Equal(t, result, string(jsonData))
	jsonBook := format.Book{}
	err = jsonBook.UnmarshalJSON(jsonData)
	assert.Nil(t, err)
	assert.Equal(t, book, jsonBook)
}

func Test_Xml(t *testing.T) {
	xmlData, err := book.WriteXML()
	assert.Nil(t, err)
	result := `<Book><id>1</id><title>Go Programming</title><author>John Doe</author><year>2021</year>` +
		`<size>350</size><rate>4.5</rate><sample>SomeBytes</sample></Book>`
	assert.Equal(t, result, string(xmlData))
	xmlBook := format.Book{}
	err = xmlBook.ReadXML(xmlData)
	assert.Nil(t, err)
	assert.Equal(t, book, xmlBook)
}

func Test_Yaml(t *testing.T) {
	yamlData, err := book.MarshalYAML()
	assert.Nil(t, err)
	result := "alias:\n  id: 1\n  title: Go Programming\n  author: John Doe\n " +
		" year: 2021\n  size: 350\n  rate: 4.5\n  sample:\n  - 83\n  - 111\n  - 109\n" +
		"  - 101\n  - 66\n  - 121\n  - 116\n  - 101\n  - 115\n"
	assert.Equal(t, result, string(yamlData))
	yamlBook := format.Book{}
	err = yamlBook.UnmarshalYAML(yamlData)
	assert.Nil(t, err)
	assert.Equal(t, book, yamlBook)
}

func Test_Gob(t *testing.T) {
	gobData, err := book.EncodeGob()
	assert.Nil(t, err)
	result := "T\xff\x81\x03\x01\x01\x04Book\x01\xff\x82\x00\x01\a\x01\x02ID" +
		"\x01\x04\x00\x01\x05Title\x01\f\x00\x01\x06Author\x01\f\x00\x01\x04Year" +
		"\x01\x04\x00\x01\x04Size\x01\x04\x00\x01\x04Rate\x01\b\x00\x01\x06Sample" +
		"\x01\n\x00\x00\x006\xff\x82\x01\x02\x01\x0eGo Programming\x01\bJohn Doe" +
		"\x01\xfe\x0f\xca\x01\xfe\x02\xbc\x01\xfe\x12@\x01\tSomeBytes\x00"
	assert.Equal(t, result, string(gobData))
	gobBook := format.Book{}
	err = gobBook.DecodeGob(gobData)
	assert.Nil(t, err)
	assert.Equal(t, book, gobBook)
}

func Test_bj(t *testing.T) {
	bjsonData, err := book.MarshalMsgpack()
	assert.Nil(t, err)
	result := "\x87\xa2ID\x01\xa5Title\xaeGo Programming\xa6Author\xa8John Doe\xa4Year" +
		"\xcd\a\xe5\xa4Size\xcd\x01^\xa4Rate\xcb@\x12\x00\x00\x00\x00\x00\x00\xa6Sample" +
		"\xc4\tSomeBytes"
	assert.Equal(t, result, string(bjsonData))
	bjsonBook := format.Book{}
	err = bjsonBook.UnmarshalMsgpack(bjsonData)
	assert.Nil(t, err)
	assert.Equal(t, book, bjsonBook)
}
