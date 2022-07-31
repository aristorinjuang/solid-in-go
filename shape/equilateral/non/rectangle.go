package non

type Rectangle struct {
	length, width float64
}

func (r *Rectangle) SetA(length float64) {
	r.length = length
}

func (r *Rectangle) SetB(width float64) {
	r.width = width
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}
