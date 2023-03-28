package shape

type Rectangle struct {
	Width  int
	Length int
}

func (rectangle *Rectangle) Area() int {
	return rectangle.Length * rectangle.Width
}

func (rectangle *Rectangle) Parameter() int {
	return 2 * (rectangle.Length + rectangle.Width)
}
