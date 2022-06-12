package premium

import "github.com/aristorinjuang/solid-in-go/user"

type PremiumUser interface {
	user.User
	ShapeCircumferences() float64
	AverageShapeCircumference() float64
}
