package pixel

import (
	"fmt"
	"math"
)

type Color struct {
	red   float64
	green float64
	blue  float64
}

func NewColor(r uint16, g uint16, b uint16) Color {
	return Color{
		red:   float64(r)/65535.0,
		green: float64(g)/65535.0,
		blue:  float64(b)/65535.0,
	}
}

func (c Color) GetLuminance() float64 {
	return float64(c.red)*0.2126 + float64(c.green)*0.7152 + float64(c.blue)*0.0722
}

func (c Color) GetRed() uint16 {
	return uint16(c.red * 65535)
}

func (c Color) GetGreen() uint16 {
	return uint16(c.green * 65535)
}

func (c Color) GetBlue() uint16 {
	return uint16(c.blue * 65535)
}

func (c Color) Add(col Color) Color {
	return Color{
		red:   math.Min(1.0, c.red+col.red),
		green: math.Min(1.0, c.green+col.green),
		blue:  math.Min(1.0, c.blue+col.blue),
	}
}

func (c Color) AddFactor(n float64) Color {
	return Color{
		red:   math.Min(1.0, c.red+n),
		green: math.Min(1.0, c.green+n),
		blue: math.Min(1.0, c.blue+n),
	}
}

func (c Color) Multiply(brightness float64) Color {
	return Color{
		red: math.Min(1.0, c.red * brightness),
		green: math.Min(1.0, c.green * brightness),
		blue: math.Min(1.0, c.blue * brightness),
	}
}

func (c Color) String() string {
	return fmt.Sprintf("{r:%v,g:%v,b:%v}", c.red, c.green, c.blue)
}

func limit(x float64) uint16 {
	if x < 0 || x > 65535 {
		return 65535
	}
	return uint16(x)
}

func Lerp(a Color, b Color, factor float64) Color {
	dr := b.red - a.red
	dg := b.green - a.green
	db := b.blue - a.blue

	return Color{
		red:   a.red + (dr * factor),
		green: a.green + (dg * factor),
		blue: a.blue + (db * factor),
	}
}
