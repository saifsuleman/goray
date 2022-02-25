package ray

import (
	"github.com/saifsuleman/goray/vec"
)

type Shape interface {
	getNormalAt(origin vec.Vector, point vec.Vector) vec.Vector
	calculateIntersection(origin *vec.Vector, ray *Ray) (vec.Vector, bool)
}
