package memory

import (
	"github.com/aristorinjuang/solid-in-go/shape"
	"github.com/aristorinjuang/solid-in-go/user"
)

type Memory struct {
	shapes              map[uint32][]shape.Shape
	shapeAreas          map[uint32]float64
	shapeCircumferences map[uint32]float64
}

func (m *Memory) AddShape(u user.User, s shape.Shape) {
	m.shapes[u.ID()] = append(m.shapes[u.ID()], s)
	m.shapeAreas[u.ID()] += s.Area()
	m.shapeCircumferences[u.ID()] += s.Circumference()
}

func (m *Memory) Shapes(u user.User) []shape.Shape {
	return m.shapes[u.ID()]
}

func (m *Memory) ShapeAreas(u user.User) float64 {
	return m.shapeAreas[u.ID()]
}

func (m *Memory) ShapeCircumferences(u user.User) float64 {
	return m.shapeCircumferences[u.ID()]
}

func NewMemory() *Memory {
	return &Memory{
		shapes:              make(map[uint32][]shape.Shape),
		shapeAreas:          make(map[uint32]float64),
		shapeCircumferences: make(map[uint32]float64),
	}
}
