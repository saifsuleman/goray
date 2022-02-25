package ray

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/jpeg"
	"math"

	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/vec"
)

//go:embed Sky.jpg
var skyboxBytes []byte

type Camera struct {
	Position vec.Vector
	Yaw      float64
	Pitch    float64
}

type Scene struct {
	Entities []*Entity
	Camera   *Camera
	Light    *LightSource
	Image    image.Image
}

func (s *Scene) GetSkyboxColor(vec vec.Vector) pixel.Color {
	u := 0.5+math.Atan2(vec.GetZ(), vec.GetX())/(2*math.Pi)
	v := 0.5 - math.Asin(vec.GetY())/math.Pi
	color := s.Image.At(int(u*(float64(s.Image.Bounds().Dx())-1.0)), int(v*(float64(s.Image.Bounds().Dy())-1.0)))
	r, g, b, _ := color.RGBA()
	return pixel.NewColor(uint16(r), uint16(g), uint16(b))
}

func NewScene() Scene {
	light := NewLightSource(vec.NewVector(10, 10, 10), 1.0)

	image, _, err := image.Decode(bytes.NewReader(skyboxBytes))
	if err != nil {
		panic(err)
	}

	return Scene{
		Entities: []*Entity{},
		Camera: &Camera{
			Position: vec.NewVector(0, 0, 0),
			Yaw:      0,
			Pitch:    0,
		},
		Light: &light,
		Image: image,
	}
}

func (s *Scene) AddEntity(entity *Entity) {
	s.Entities = append(s.Entities, entity)
}
