package ray

import (
	"math"

	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

type Box struct {
	Scale vec.Vector
	Color pixel.Color
}

func (b *Box) calculateIntersection(origin vec.Vector, ray *Ray) (vec.Vector, bool) {
	min := origin.SubtractVector(b.Scale.Multiply(0.5))
	max := origin.AddVector(b.Scale.Multiply(0.5))

	var t1 float64
	var t2 float64
	tnear := math.Inf(-1)
	tfar := math.Inf(1)
	intersectFlag := true

	rayDirection := ray.Direction.ToArray()
	rayOrigin := ray.Origin.ToArray()

	b1 := min.ToArray()
	b2 := max.ToArray()

	for i := 0; i < 3; i++ {
		if rayDirection[i] == 0 {
			if rayOrigin[i] < b1[i] || rayOrigin[i] > b2[i] {
				intersectFlag = false
			}
		} else {
			t1 = (b1[i] - rayOrigin[i]) / rayDirection[i]
			t2 = (b2[i] - rayOrigin[i]) / rayDirection[i]
			if t1 > t2 {
				t1, t2 = t2, t1
			}
			if t1 > tnear {
				tnear = t1
			}
			if t2 < tfar {
				tfar = t2
			}
			if tnear > tfar {
				intersectFlag = false
			}
			if tfar < 0 {
				intersectFlag = false
			}
		}
	}

	if !intersectFlag {
		return vec.Vector{}, false
	}

	return ray.Origin.AddVector(ray.Direction.Multiply(tnear)), true
}

func (b *Box) getNormalAt(origin vec.Vector, point vec.Vector) vec.Vector {
	direction := point.SubtractVector(origin).ToArray()
	biggest := 0.0

	for i := 0; i < 3; i++ {
		if biggest < math.Abs(direction[i]) {
			biggest = math.Abs(direction[i])
		}
	}

	if biggest == 0 {
		return vec.NewVector(0, 0, 0)
	}

	for i := 0; i < 3; i++ {
		if math.Abs(direction[i]) == biggest {
			normal := []float64{0, 0, 0}
			normal[i] = direction[i] / math.Abs(direction[i])
			return vec.NewVector(normal[0], normal[1], normal[2])
		}
	}

	return vec.NewVector(0, 0, 0)
}

func (b *Box) GetColorAt(point vec.Vector) pixel.Color {
	return b.Color
}

func NewBox(position vec.Vector, dimensions vec.Vector, color pixel.Color) *Entity {
	box := &Box{Scale: dimensions, Color: color}
	return &Entity{
		Position:      position,
		Shape:         box,
		Reflectivity:  0.4,
		Emission:      0.4,
		ColorProvider: box,
	}
}
