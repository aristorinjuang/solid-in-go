package non

import "github.com/aristorinjuang/solid-in-go/shape/equilateral"

type NonEquilateral interface {
	equilateral.Equilateral
	SetB(float64)
}
