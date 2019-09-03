package aseprite

import "image"

// WidthHeight is a basic object to store width and height.
type WidthHeight struct {
	Width  int `json:"w"`
	Height int `json:"h"`
}

// Boundary is a basic object to store position coordinates, width, and height.
type Boundary struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
}

// Rectangle returns a basic rectangle of the belonging coordinates.
func (b Boundary) Rectangle() image.Rectangle {
	return image.Rect(
		b.X, b.Y,
		b.X+b.Width, b.Y+b.Height,
	)
}
