package format

import (
	"encoding/xml"
)

func (b Book) WriteXML() ([]byte, error) {
	return xml.Marshal(b)
}

func (b *Book) ReadXML(data []byte) error {
	type Alias Book
	aux := &struct {
		Alias
	}{Alias: (Alias)(*b)}
	err := xml.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	*b = Book(aux.Alias)
	return nil
}
