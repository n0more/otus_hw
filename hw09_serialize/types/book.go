package types

type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title"`
	Author string  `json:"author" xml:"author" yaml:"author"`
	Year   int     `json:"year" xml:"year" yaml:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample"`
}
