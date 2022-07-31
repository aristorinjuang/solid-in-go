package premium

import "github.com/aristorinjuang/solid-in-go/user"

type PremiumUser interface {
	user.User
	ShapePerimeters() float64
	AverageShapePerimeter() float64
}
