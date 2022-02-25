package vec

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func NewVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func DotVector(a Vector, b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (v Vector) GetX() float64 {
	return v.x
}

func (v Vector) GetY() float64 {
	return v.y
}

func (v Vector) GetZ() float64 {
	return v.z
}

func (v Vector) Add(x float64, y float64, z float64) Vector {
	return NewVector(v.x+x, v.y+y, v.z+z)
}

func (v Vector) AddVector(vec Vector) Vector {
	return v.Add(vec.x, vec.y, vec.z)
}

func (v Vector) Subtract(x float64, y float64, z float64) Vector {
	return NewVector(v.x-x, v.y-y, v.z-z)
}

func (v Vector) SubtractVector(vec Vector) Vector {
	return v.Subtract(vec.x, vec.y, vec.z)
}

func (v Vector) Multiply(n float64) Vector {
	return NewVector(v.x*n, v.y*n, v.z*n)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return NewVector(v.x/l, v.y/l, v.z/l)
}

func (v Vector) RotateYP(yaw float64, pitch float64) Vector {
	yawRads := yaw * (math.Pi / 180)
	pitchRads := pitch * (math.Pi / 180)
	y := v.y*math.Cos(pitchRads) - v.z*math.Sin(pitchRads)
	z := v.y*math.Sin(pitchRads) + v.z*math.Cos(pitchRads)
	x := v.x*math.Cos(yawRads) + z*math.Sin(yawRads)
	z = -v.x*math.Sin(yawRads) + z*math.Cos(yawRads)
	return NewVector(x, y, z)
}

func (v Vector) Distance(vec Vector) float64 {
	dx := vec.x - v.x
	dy := vec.y - v.y
	dz := vec.z - v.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
