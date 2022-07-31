package memory

import (
	"github.com/aristorinjuang/solid-in-go/shape"
	"github.com/aristorinjuang/solid-in-go/user"
)

type Memory struct {
	shapes          map[uint32][]shape.Shape
	shapeAreas      map[uint32]float64
	shapePerimeters map[uint32]float64
}

func (m *Memory) AddShape(u user.User, s shape.Shape) {
	m.shapes[u.ID()] = append(m.shapes[u.ID()], s)
	m.shapeAreas[u.ID()] += s.Area()
	m.shapePerimeters[u.ID()] += s.Perimeter()
}

func (m *Memory) Shapes(u user.User) []shape.Shape {
	return m.shapes[u.ID()]
}

func (m *Memory) ShapeAreas(u user.User) float64 {
	return m.shapeAreas[u.ID()]
}

func (m *Memory) ShapePerimeters(u user.User) float64 {
	return m.shapePerimeters[u.ID()]
}

func NewMemory() *Memory {
	return &Memory{
		shapes:          make(map[uint32][]shape.Shape),
		shapeAreas:      make(map[uint32]float64),
		shapePerimeters: make(map[uint32]float64),
	}
}
