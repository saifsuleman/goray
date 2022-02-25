package pixel

type Coordinate struct {
	X int
	Y int
}

type PixelBuffer struct {
	buffer map[Coordinate]Pixel
}

type Pixel struct {
	Color    Color
	Emission float64
}

func NewBuffer() PixelBuffer {
	return PixelBuffer{
		buffer: map[Coordinate]Pixel{},
	}
}

func (pb *PixelBuffer) Set(x int, y int, pixel Pixel) {
	coord := Coordinate{X: x, Y: y}
	pb.buffer[coord] = pixel
}

func (pb *PixelBuffer) GetPixel(x int, y int) Pixel {
	return pb.buffer[Coordinate{X: x, Y: y}]
}
