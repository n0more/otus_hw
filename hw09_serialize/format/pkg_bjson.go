package format

import (
	"github.com/vmihailenco/msgpack/v5"
)

func (b Book) MarshalMsgpack() ([]byte, error) {
	type Alias Book
	return msgpack.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

func (b *Book) UnmarshalMsgpack(data []byte) error {
	type Alias Book
	aux := &struct {
		Alias
	}{Alias: (Alias)(*b)}
	err := msgpack.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	*b = Book(aux.Alias)
	return nil
}
