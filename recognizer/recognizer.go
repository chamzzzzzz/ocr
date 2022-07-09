package recognizer

type Image struct {
	File   string
	Width  int
	Height int
}

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

type BoudingBox struct {
	Origin Point
	Size   Size
}

type Observation struct {
	Confidence int
	Text       string
	BoudingBox BoudingBox
}

type Result struct {
	Image        Image
	Observations []*Observation
}

type Recognizer interface {
	Recognize(file string) (*Result, error)
}
