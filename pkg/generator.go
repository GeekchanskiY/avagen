package pkg

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Scene struct {
	Width, Height int
	Img           *image.RGBA
}

func NewScene(width int, height int) *Scene {
	return &Scene{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func randomColor() color.RGBA {
	src := rand.New(rand.NewSource(time.Now().Unix()))

	return color.RGBA{
		R: uint8(src.Intn(255)),
		G: uint8(src.Intn(255)),
		B: uint8(src.Intn(255)),
		A: 255}
}

func (s *Scene) setPixel(x int, y int, color color.RGBA) {
	s.Img.Set(x, y, color)
}

func generateImage(w, h int, pixelColor color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			img.Set(x, y, pixelColor)
		}
	}
	return img
}

func (s *Scene) EachPixel(colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.setPixel(x, y, colorFunction(x, y))
		}
	}
}

func (s *Scene) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, s.Img)
}
