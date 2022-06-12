package user

import "github.com/aristorinjuang/solid-in-go/shape"

type User interface {
	ID() uint32
	AddShape(shape.Shape)
	ShapeAreas() float64
	AverageShapeArea() float64
}
