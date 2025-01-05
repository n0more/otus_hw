package main

import (
	"fmt"

	"github.com/n0more/otus_hw/hw09_serialize/format"
)

func main() {
	book := format.Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2021,
		Size:   350,
		Rate:   4.5,
		Sample: []byte("SomeBytes"),
	}
	fmt.Println("Book:", book)

	jsonData, err := book.MarshalJSON()
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	jsonBook := format.Book{}

	err = jsonBook.UnmarshalJSON(jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling from JSON:", err)
		return
	}

	fmt.Println("Book:", jsonBook)

	jsonData2, err := jsonBook.MarshalJSON()
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData2))

	xmlData, err := book.WriteXML()
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	fmt.Println("XML:", string(xmlData))

	xmlBook := format.Book{}

	err = xmlBook.ReadXML(xmlData)
	if err != nil {
		fmt.Println("Error unmarshaling from XML:", err)
		return
	}
	fmt.Println("Book:", xmlBook)

	xmlData2, err := xmlBook.WriteXML()
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	fmt.Println("XML:", string(xmlData2))

	yamlData, err := book.MarshalYAML()
	if err != nil {
		fmt.Println("Error marshaling to YAML:", err)
		return
	}
	fmt.Println("YAML:", string(yamlData))

	yamlBook := format.Book{}
	err = yamlBook.UnmarshalYAML(yamlData)
	if err != nil {
		fmt.Println("Error unmarshaling from YAML:", err)
		return
	}
	fmt.Println("Book:", yamlBook)
	yamlData2, err := yamlBook.MarshalYAML()
	if err != nil {
		fmt.Println("Error marshaling to YAML:", err)
		return
	}
	fmt.Println("YAML:", string(yamlData2))
}
