package premium

import (
	"github.com/aristorinjuang/solid-in-go/shape"
	"github.com/aristorinjuang/solid-in-go/user"
)

type Member struct {
	user.Guest
}

func (m *Member) AddShape(s shape.Shape) {
	if !user.HasPermission(m, s) {
		return
	}

	m.Repo.AddShape(m, s)
}

func (m *Member) ShapeCircumferences() float64 {
	return m.Repo.ShapeCircumferences(m)
}

func (m *Member) AverageShapeCircumference() float64 {
	return m.Repo.ShapeCircumferences(m) / float64(len(m.Repo.Shapes(m)))
}

func NewMember(r user.Repository) *Member {
	return &Member{*user.NewGuest(r)}
}
