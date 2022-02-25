package ray

import (
	"image"
	"math"

	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

type Sphere struct {
	Radius  float64
	Color   pixel.Color
	Texture image.Image
}

func (s *Sphere) getNormalAt(origin vec.Vector, point vec.Vector) vec.Vector {
	vec := point.Subtract(origin.GetX(), origin.GetY(), origin.GetZ())
	return vec.Normalize()
}

func (s *Sphere) calculateIntersection(origin vec.Vector, ray *Ray) (vec.Vector, bool) {
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

func (s *Sphere) GetColorAt(vec vec.Vector) pixel.Color {
	if s.Texture != nil {
		u := 0.5 + math.Atan2(vec.GetZ(), vec.GetX())/(2*math.Pi)
		v := 0.5 - math.Asin(vec.GetY())/math.Pi
		color := s.Texture.At(int(u*(float64(s.Texture.Bounds().Dx())-1.0)), int(v*(float64(s.Texture.Bounds().Dy())-1.0)))
		r, g, b, _ := color.RGBA()
		return pixel.NewColor(uint16(r), uint16(g), uint16(b))
	}

	return s.Color
}

func NewTexturedSphere(position vec.Vector, radius float64, img image.Image) *Entity {
	sphere := &Sphere{Radius: radius, Texture: img}
	return &Entity{
		Position:      position,
		Shape:         sphere,
		Reflectivity:  0.4,
		Emission:      0.4,
		ColorProvider: sphere,
	}
}

func NewSphere(position vec.Vector, radius float64, color pixel.Color) *Entity {
	sphere := &Sphere{Radius: radius, Color: color}
	return &Entity{
		Position:      position,
		Shape:         sphere,
		Reflectivity:  0.4,
		Emission:      0.4,
		ColorProvider: sphere,
	}
}
