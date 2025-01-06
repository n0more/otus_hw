package format

import (
	"encoding/json"
)

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		Alias
	}{Alias: (Alias)(*b)}
	err := json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	*b = Book(aux.Alias)
	return nil
}
