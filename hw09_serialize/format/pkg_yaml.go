package format

import (
	"gopkg.in/yaml.v2"
)

func (b Book) MarshalYAML() ([]byte, error) {
	type Alias Book
	return yaml.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(b),
	})
}

func (b *Book) UnmarshalYAML(data []byte) error {
	type Alias Book
	aux := &struct {
		Alias
	}{Alias: (Alias)(*b)}
	if err := yaml.Unmarshal(data, aux); err != nil {
		return err
	}

	*b = Book(aux.Alias)
	return nil
}
