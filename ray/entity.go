package ray

import (
	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

type Entity struct {
	Position      vec.Vector
	Shape         Shape
	ColorProvider ColorProvider
	Reflectivity  float64
	Emission      float64
}

type ColorProvider interface {
	GetColorAt(vec.Vector) pixel.Color
}

func (e *Entity) GetNormalAt(point vec.Vector) vec.Vector {
	return e.Shape.getNormalAt(e.Position, point)
}

func (e *Entity) CalculateIntersection(ray *Ray) (vec.Vector, bool) {
	return e.Shape.calculateIntersection(e.Position, ray)
}
