package format

import (
	"bytes"
	"encoding/gob"
)

func (b Book) EncodeGob() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Book) DecodeGob(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}
