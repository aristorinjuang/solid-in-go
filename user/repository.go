package user

import (
	"github.com/aristorinjuang/solid-in-go/shape"
)

type Repository interface {
	AddShape(User, shape.Shape)
	Shapes(User) []shape.Shape
	ShapeAreas(User) float64
	ShapePerimeters(User) float64
}
