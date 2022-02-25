package main

import (
	"bytes"
	"fmt"
	"image"

	_ "embed"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/ray"
	"github.com/saifsuleman/goray/renderer"
	"github.com/saifsuleman/goray/vec"
)

//go:embed download.jpeg
var grass []byte

type Game struct {
	Scene      *ray.Scene
	Resolution float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0, 0, 1).RotateYP(g.Scene.Camera.Yaw, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0, 0, -1).RotateYP(g.Scene.Camera.Yaw, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(1, 0, 0).RotateYP(g.Scene.Camera.Yaw, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(-1, 0, 0).RotateYP(g.Scene.Camera.Yaw, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0, 1, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0, -1, 0))
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Scene.Camera.Yaw += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Scene.Camera.Yaw -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Scene.Camera.Pitch += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Scene.Camera.Pitch -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Resolution = 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyT) {
		g.Resolution = 1.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		go renderer.RenderFile("output.jpg", g.Scene, 1920, 1080)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	renderer.Render(g.Scene, 1280, 720, g.Resolution, screen)

	ebiten.SetWindowTitle(fmt.Sprintf("Raytracing - FPS: %v", ebiten.CurrentFPS()))
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 720
}

func main() {
	scene := ray.NewScene()
	img, _, err := image.Decode(bytes.NewReader(grass))
	if err != nil {
		panic(err)
	}
	scene.AddEntity(ray.NewSphere(vec.NewVector(-8, 0, 10), 3, pixel.NewColor(65535, 0, 0)))
	scene.AddEntity(ray.NewTexturedSphere(vec.NewVector(0, 0, 10), 3, img))
	// scene.AddEntity(ray.NewSphere(vec.NewVector(15, 0, 10), 3, pixel.NewColor(65535, 17321, 12444)))
	scene.AddEntity(ray.NewBox(vec.NewVector(15, 0, 10), vec.NewVector(3, 3, 3), pixel.NewColor(65535, 32768, 0)))
	scene.AddEntity(ray.NewPlane())
	mirror := ray.NewSphere(vec.NewVector(8, 0, 10), 3, pixel.NewColor(115*257, 245*257, 139*257))
	mirror.Reflectivity = 1.0
	mirror.Emission = 0.0
	scene.AddEntity(mirror)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Raytracing")

	game := &Game{Scene: &scene, Resolution: 1}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
