package ray

import "github.com/saifsuleman/goray/vec"

type LightSource struct {
	position  vec.Vector
	intensity float64
}

func NewLightSource(position vec.Vector, intensity float64) LightSource {
	return LightSource{position, intensity}
}

func (ls *LightSource) GetPosition() vec.Vector {
	return ls.position
}

func (ls *LightSource) GetIntensity() float64 {
	return ls.intensity
}
