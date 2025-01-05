package main

import (
	"fmt"

	"github.com/n0more/otus_hw/hw09_serialize/format"
)

func tryJSON(book format.Book) {
	jsonData, err := book.MarshalJSON()
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	jsonBook := format.Book{}
	if unmarshalErr := jsonBook.UnmarshalJSON(jsonData); unmarshalErr != nil {
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
}

func tryXML(book format.Book) {
	xmlData, err := book.WriteXML()
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	fmt.Println("XML:", string(xmlData))

	xmlBook := format.Book{}
	if readErr := xmlBook.ReadXML(xmlData); readErr != nil {
		fmt.Println("Error unmarshaling from XML:", readErr)
		return
	}
	fmt.Println("Book:", xmlBook)

	xmlData2, err := xmlBook.WriteXML()
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	fmt.Println("XML:", string(xmlData2))
}

func tryYAML(book format.Book) {
	yamlData, err := book.MarshalYAML()
	if err != nil {
		fmt.Println("Error marshaling to YAML:", err)
		return
	}
	fmt.Println("YAML:", string(yamlData))

	yamlBook := format.Book{}
	if unmarshalErr := yamlBook.UnmarshalYAML(yamlData); unmarshalErr != nil {
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

func tryGOB(book format.Book) {
	gobData, err := book.EncodeGob()
	if err != nil {
		fmt.Println("Error marshaling to GOB:", err)
		return
	}
	fmt.Println("GOB:", gobData)

	gobBook := format.Book{}
	if decodeErr := gobBook.DecodeGob(gobData); decodeErr != nil {
		fmt.Println("Error unmarshaling from GOB:", err)
		return
	}
	fmt.Println("Book:", gobBook)

	gobData2, err := gobBook.EncodeGob()
	if err != nil {
		fmt.Println("Error marshaling to GOB:", err)
		return
	}
	fmt.Println("GOB:", gobData2)
}

func tryMSGPACK(book format.Book) {
	bjsonData, err := book.MarshalMsgpack()
	if err != nil {
		fmt.Println("Error marshaling to MSGPACK:", err)
		return
	}
	fmt.Println("MSGPACK:", bjsonData)

	bjsonBook := format.Book{}
	if unmarshalErr := bjsonBook.UnmarshalMsgpack(bjsonData); unmarshalErr != nil {
		fmt.Println("Error unmarshaling from MSGPACK:", err)
		return
	}
	fmt.Println("Book:", bjsonBook)

	bjsonData2, err := bjsonBook.MarshalMsgpack()
	if err != nil {
		fmt.Println("Error marshaling to MSGPACK:", err)
		return
	}
	fmt.Println("MSGPACK:", bjsonData2)
}

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

	tryJSON(book)
	tryXML(book)
	tryYAML(book)
	tryGOB(book)
	tryMSGPACK(book)
}
