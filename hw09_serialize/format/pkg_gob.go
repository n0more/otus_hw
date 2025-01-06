package format

import (
	"bytes"
	"encoding/gob"
)

func (b Book) EncodeGob() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Book) DecodeGob(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(b)
	if err != nil {
		return err
	}
	return nil
}
