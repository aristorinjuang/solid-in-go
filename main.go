package main

import (
	"fmt"

	"github.com/aristorinjuang/solid-in-go/shape/equilateral"
	"github.com/aristorinjuang/solid-in-go/shape/equilateral/non"
	"github.com/aristorinjuang/solid-in-go/user"
	"github.com/aristorinjuang/solid-in-go/user/memory"
	"github.com/aristorinjuang/solid-in-go/user/premium"
)

func main() {
	var s1 equilateral.Equilateral = new(equilateral.Circle)
	s1.SetA(5)
	fmt.Println("s1.Area()", s1.Area())

	var s2 equilateral.Equilateral = new(equilateral.Square)
	s2.SetA(4)
	fmt.Println("s2.Area()", s2.Area())

	var db user.Repository = memory.NewMemory()

	var u1 user.User = user.NewGuest(db)
	u1.AddShape(s1)
	u1.AddShape(s2)

	fmt.Println("u1.AverageShapeArea()", u1.AverageShapeArea())
	fmt.Println("u1.ShapeAreas()", u1.ShapeAreas())

	var s3 non.NonEquilateral = new(non.Rectangle)
	s3.SetA(6)
	s3.SetB(7)
	fmt.Println("s3.Area()", s3.Area())
	fmt.Println("s3.Perimeter()", s3.Perimeter())

	var u2 premium.PremiumUser = premium.NewMember(db)
	u2.AddShape(s1)
	u2.AddShape(s2)
	u2.AddShape(s3)

	fmt.Println("u2.AverageShapeArea()", u2.AverageShapeArea())
	fmt.Println("u2.ShapeAreas()", u2.ShapeAreas())
	fmt.Println("u2.AverageShapePerimeter()", u2.AverageShapePerimeter())
	fmt.Println("u2.ShapePerimeters()", u2.ShapePerimeters())
}
