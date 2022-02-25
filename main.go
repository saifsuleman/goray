package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/saifsuleman/goray/pixel"
	"github.com/saifsuleman/goray/ray"
	"github.com/saifsuleman/goray/renderer"
	"github.com/saifsuleman/goray/vec"
	"image/color"
)

type Game struct {
	Scene *ray.Scene
	Resolution float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0,0, 2).RotateYP(g.Scene.Camera.Yaw, g.Scene.Camera.Pitch))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(0,0, -2).RotateYP(g.Scene.Camera.Yaw, g.Scene.Camera.Pitch))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(2,0, 0).RotateYP(g.Scene.Camera.Yaw, g.Scene.Camera.Pitch))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Scene.Camera.Position = g.Scene.Camera.Position.AddVector(vec.NewVector(-2,0, 0).RotateYP(g.Scene.Camera.Yaw, g.Scene.Camera.Pitch))
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
		if g.Resolution == 1.0 {
			g.Resolution = 0.1
		} else {
			g.Resolution = 1.0
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	buf := renderer.Render(g.Scene, 1280, 720, g.Resolution)

	for x := 0; x < 1280; x++ {
		for y := 0; y < 720; y++ {
			pixelColor := buf.GetPixel(x, y).Color
			screen.Set(x, y, color.RGBA64{
				R: pixelColor.GetRed(),
				G: pixelColor.GetGreen(),
				B: pixelColor.GetBlue(),
			})
		}
	}

	ebiten.SetWindowTitle(fmt.Sprintf("Raytracing - FPS: %v", ebiten.CurrentFPS()))
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 720
}

func main() {
	scene := ray.NewScene()
	scene.AddEntity(ray.NewSphere(vec.NewVector(0, 0, 10), 3, pixel.NewColor(65535, 0, 0)))
	scene.AddEntity(ray.NewSphere(vec.NewVector(8, 0, 10), 3, pixel.NewColor(235*257,212*257,48*257)))

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Raytracing")

	game := &Game{Scene: &scene,Resolution: 0.1}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
