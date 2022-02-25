package ray

import (
	"math"

	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

type Plane struct{}

func (p *Plane) getNormalAt(origin vec.Vector, point vec.Vector) vec.Vector {
	return vec.NewVector(0, 1, 0)
}

func (p *Plane) calculateIntersection(origin vec.Vector, ray *Ray) (vec.Vector, bool) {
	t := -(ray.Origin.GetY() - origin.GetY()) / ray.Direction.GetY()
	if t <= 0 || t > 1000 {
		return vec.Vector{}, false
	}
	return ray.Origin.AddVector(ray.Direction.Multiply(t)), true
}

func (p *Plane) GetColorAt(point vec.Vector) pixel.Color {
	BLACK := pixel.NewColor(5000, 5000, 5000)
	WHITE := pixel.NewColor(13364, 13364, 13364)

	x := int(math.Abs(math.Floor(point.GetX() / 10)))
	z := int(math.Abs(math.Floor(point.GetZ() / 10)))

	if x%2 != z%2 {
		return WHITE
	} else {
		return BLACK
	}
}

func NewPlane() *Entity {
	plane := &Plane{}
	return &Entity{
		Position:      vec.NewVector(0, -5, 0),
		Shape:         plane,
		Reflectivity:  0.1,
		Emission:      1.0,
		ColorProvider: plane,
	}
}
