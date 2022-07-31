package premium

import (
	"github.com/aristorinjuang/solid-in-go/shape"
	"github.com/aristorinjuang/solid-in-go/user"
)

type Member struct {
	user.Guest
}

func (m *Member) AddShape(s shape.Shape) {
	m.Repo.AddShape(m, s)
}

func (m *Member) ShapePerimeters() float64 {
	return m.Repo.ShapePerimeters(m)
}

func (m *Member) AverageShapePerimeter() float64 {
	return m.Repo.ShapePerimeters(m) / float64(len(m.Repo.Shapes(m)))
}

func NewMember(r user.Repository) *Member {
	return &Member{*user.NewGuest(r)}
}
