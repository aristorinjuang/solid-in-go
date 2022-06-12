package equilateral

import "math"

type Square struct {
	side float64
}

func (s *Square) SetA(side float64) {
	s.side = side
}

func (s *Square) Area() float64 {
	return math.Pow(s.side, 2)
}

func (s *Square) Circumference() float64 {
	return 4 * s.side
}
