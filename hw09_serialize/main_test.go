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
	jsonBook := format.Book{}
	err = jsonBook.UnmarshalJSON(jsonData)
	assert.Nil(t, err)
	assert.Equal(t, book, jsonBook)
}

func Test_Xml(t *testing.T) {
	xmlData, err := book.WriteXML()
	assert.Nil(t, err)
	xmlBook := format.Book{}
	err = xmlBook.ReadXML(xmlData)
	assert.Nil(t, err)
	assert.Equal(t, book, xmlBook)
}

func Test_Yaml(t *testing.T) {
	yamlData, err := book.MarshalYAML()
	assert.Nil(t, err)
	yamlBook := format.Book{}
	err = yamlBook.UnmarshalYAML(yamlData)
	assert.Nil(t, err)
	assert.Equal(t, book, yamlBook)
}

func Test_Gob(t *testing.T) {
	gobData, err := book.EncodeGob()
	assert.Nil(t, err)
	gobBook := format.Book{}
	err = gobBook.DecodeGob(gobData)
	assert.Nil(t, err)
	assert.Equal(t, book, gobBook)
}

func Test_bj(t *testing.T) {
	bjsonData, err := book.MarshalMsgpack()
	assert.Nil(t, err)
	bjsonBook := format.Book{}
	err = bjsonBook.UnmarshalMsgpack(bjsonData)
	assert.Nil(t, err)
	assert.Equal(t, book, bjsonBook)
}
