package ray

import (
	"math"

	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) getNormalAt(origin vec.Vector, point vec.Vector) vec.Vector {
	vec := point.Subtract(origin.GetX(), origin.GetY(), origin.GetZ())
	return vec.Normalize()
}

func (s *Sphere) calculateIntersection(origin *vec.Vector, ray *Ray) (vec.Vector, bool) {
	t := vec.DotVector(origin.SubtractVector(ray.Origin), ray.Direction)
	p := ray.Origin.AddVector(ray.Direction.Multiply(t))
	yv := origin.SubtractVector(p)
	y := yv.Length()

	if y < s.Radius {
		x := math.Sqrt(s.Radius*s.Radius - y*y)
		t1 := t - x
		if t1 > 0 {
			return ray.Origin.AddVector(ray.Direction.Multiply(t1)), true
		}
	}
	return vec.Vector{}, false
}

func NewSphere(position vec.Vector, radius float64, color pixel.Color) *Entity {
	sphere := &Sphere{Radius: radius}
	return &Entity{
		Position:     &position,
		Shape:        sphere,
		Reflectivity: 0.4,
		Emission:     0.6,
		Color:        color,
	}
}
