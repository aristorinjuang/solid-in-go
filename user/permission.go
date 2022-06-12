package user

import (
	"fmt"
	"strings"

	"github.com/aristorinjuang/solid-in-go/shape"
)

func HasPermission(u User, s shape.Shape) bool {
	if fmt.Sprintf("%T", u) != "*user.Guest" {
		return true
	}
	if strings.Contains(fmt.Sprintf("%T", s), "*equilateral.") {
		return true
	}
	return false
}
