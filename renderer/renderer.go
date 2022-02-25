package renderer

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	"github.com/skratchdot/open-golang/open"

	"github.com/hajimehoshi/ebiten"
	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/ray"
	"github.com/saifsuleman/goray/vec"
)

const (
	GLOBAL_ILLUMINATION    = 0.3
	SKY_EMISSION           = 0.5
	MAX_REFLECTION_BOUNCES = 4
)

func Render(scene *ray.Scene, width int, height int, resolution float64, screen *ebiten.Image) *pixel.PixelBuffer {
	buf := pixel.NewBuffer()

	blockSize := int(1.0 / resolution)

	for x := 0; x < width; x += blockSize {
		for y := 0; y < height; y += blockSize {
			u, v := getNormalizedScreenCoordinates(x, y, width, height)
			p := computePixel(scene, u, v)

			for i := 0; i < blockSize; i++ {
				for j := 0; j < blockSize; j++ {
					screen.Set(x+i, y+j, color.RGBA64{
						R: p.Color.GetRed(),
						G: p.Color.GetGreen(),
						B: p.Color.GetBlue(),
						A: 65535,
					})
				}
			}
		}
	}

	return &buf
}

func RenderFile(filename string, scene *ray.Scene, width int, height int) {
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			u, v := getNormalizedScreenCoordinates(x, y, width, height)
			p := computePixel(scene, u, v)
			col := color.RGBA64{
				R: p.Color.GetRed(),
				G: p.Color.GetGreen(),
				B: p.Color.GetBlue(),
				A: 65535,
			}
			img.Set(x, y, col)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		panic(err)
	}

	open.Start(filename)
}

func getNormalizedScreenCoordinates(x int, y int, width int, height int) (float64, float64) {
	if width > height {
		u := (float64(x)-float64(width)/2.0+float64(height)/2.0)/float64(height)*2.0 - 1.0
		v := -(float64(y)/float64(height)*2.0 - 1.0)
		return u, v
	}

	u := float64(x)/float64(width)*2.0 - 1.0
	v := -((float64(y)-float64(height)/2+float64(width)/2)/float64(width)*2.0 - 1.0)
	return u, v
}

func computePixel(scene *ray.Scene, u float64, v float64) pixel.Pixel {
	eyePos := vec.NewVector(0, 0, -1/math.Tan((40/2)*(math.Pi/180)))
	rayDirection := vec.NewVector(u, v, 0).SubtractVector(eyePos).Normalize().RotateYP(scene.Camera.Yaw, scene.Camera.Pitch)
	r := ray.NewRay(eyePos.AddVector(scene.Camera.Position), rayDirection)
	hit := r.Cast(scene)
	if hit == nil {
		// return skybox
		return pixel.Pixel{
			Color:    scene.GetSkyboxColor(rayDirection),
			Emission: SKY_EMISSION,
		}
	}
	return computePixelHit(scene, hit, MAX_REFLECTION_BOUNCES)
}

func computePixelHit(scene *ray.Scene, hit *ray.RayHit, bounces int) pixel.Pixel {
	brightness := getDiffuseBrightness(scene, hit)
	specularBrightness := getSpecularBrightness(scene, hit)
	reflectivity := hit.Entity.Reflectivity
	emission := hit.Entity.Emission

	var reflectionHit *ray.RayHit = nil
	reflectionVector := hit.Ray.Direction.SubtractVector(hit.Normal.Multiply(2 * vec.DotVector(hit.Ray.Direction, hit.Normal)))
	reflectionOrigin := hit.Position.AddVector(reflectionVector.Multiply(0.001))
	reflectionRay := ray.NewRay(reflectionOrigin, reflectionVector)
	if bounces > 0 {
		reflectionHit = reflectionRay.Cast(scene)
	}

	var reflection pixel.Pixel
	if reflectionHit != nil {
		reflection = computePixelHit(scene, reflectionHit, bounces-1)
	} else {
		sbColor := scene.GetSkyboxColor(reflectionVector)
		reflection = pixel.Pixel{
			Color:    sbColor,
			Emission: SKY_EMISSION * sbColor.GetLuminance(),
		}
	}

	color := hit.Entity.ColorProvider.GetColorAt(hit.Position)

	pixelColor := pixel.Lerp(color, reflection.Color, reflectivity).Multiply(brightness).AddFactor(specularBrightness).
		Add(color.Multiply(emission)).
		Add(reflection.Color.Multiply(reflection.Emission * reflectivity))

	return pixel.Pixel{
		Color:    pixelColor,
		Emission: math.Min(1, emission+reflection.Emission*reflectivity+specularBrightness),
	}
}

func getDiffuseBrightness(scene *ray.Scene, hit *ray.RayHit) float64 {
	light := scene.Light
	lightRay := ray.NewRay(light.GetPosition(), hit.Position.SubtractVector(light.GetPosition()).Normalize())
	lightHit := lightRay.Cast(scene)
	if lightHit != nil && lightHit.Entity.Shape != hit.Entity.Shape {
		return GLOBAL_ILLUMINATION
	} else {
		return math.Max(GLOBAL_ILLUMINATION, math.Min(1, vec.DotVector(hit.Normal, light.GetPosition().SubtractVector(hit.Position))))
	}
}

func getSpecularBrightness(scene *ray.Scene, hit *ray.RayHit) float64 {
	cameraDirection := scene.Camera.Position.SubtractVector(hit.Position).Normalize()
	lightDirection := hit.Position.SubtractVector(scene.Light.GetPosition()).Normalize()
	lightReflectionVector := lightDirection.SubtractVector(hit.Normal.Multiply(2 * vec.DotVector(lightDirection, hit.Normal)))
	specularFactor := math.Max(0, math.Min(1, vec.DotVector(lightReflectionVector, cameraDirection)))
	return specularFactor * specularFactor * hit.Entity.Reflectivity
}
