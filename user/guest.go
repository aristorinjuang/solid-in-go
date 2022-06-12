package user

import (
	"github.com/google/uuid"

	"github.com/aristorinjuang/solid-in-go/shape"
)

type Guest struct {
	id   uuid.UUID
	Repo Repository
}

func (g *Guest) ID() uint32 {
	return g.id.ID()
}

func (g *Guest) AddShape(s shape.Shape) {
	if !HasPermission(g, s) {
		return
	}

	g.Repo.AddShape(g, s)
}

func (g *Guest) ShapeAreas() float64 {
	return g.Repo.ShapeAreas(g)
}

func (g *Guest) AverageShapeArea() float64 {
	return g.Repo.ShapeAreas(g) / float64(len(g.Repo.Shapes(g)))
}

func NewGuest(r Repository) *Guest {
	return &Guest{
		id:   uuid.New(),
		Repo: r,
	}
}
