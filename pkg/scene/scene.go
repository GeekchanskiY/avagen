package scene

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Scene struct {
	Width, Height    int
	BgColor, FgColor color.RGBA
	Img              *image.RGBA
}

func NewScene(width int, height int, BgColor, FgColor color.RGBA) *Scene {
	return &Scene{
		Width:   width,
		Height:  height,
		BgColor: BgColor,
		FgColor: FgColor,

		Img: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (s *Scene) Clear() {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.SetPixel(x, y, &s.BgColor)
		}
	}
}

func (s *Scene) EachPixel(colorFunction func(int, int) *color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.SetPixel(x, y, colorFunction(x, y))
		}
	}
}

func (s *Scene) SetPixel(x int, y int, pixelColor *color.RGBA) {
	if pixelColor == nil {
		pixelColor = &s.FgColor
	}

	s.Img.Set(x, y, *pixelColor)
}

func (s *Scene) Save(filename string) error {
	var (
		f   *os.File
		err error
	)

	if f, err = os.Create(filename); err != nil {
		return err
	}

	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	if err = png.Encode(f, s.Img); err != nil {
		return err
	}

	return nil
}
