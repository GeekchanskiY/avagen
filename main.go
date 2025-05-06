package main

import (
	"image/color"
	"log"
	"time"

	"github.com/GeekchanskiY/avagen/pkg/generators/standard"
	"github.com/GeekchanskiY/avagen/pkg/scene"
)

func main() {
	var (
		start time.Time

		width  = 400
		height = 400
		image  *scene.Scene

		err error
	)

	start = time.Now()

	image = scene.NewScene(
		width,
		height,
		color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		},
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)

	image.Clear()

	if err = standard.NewGenerator().Generate(image); err != nil {
		panic(err)
	}

	log.Printf("Time passed to generate: %s", time.Since(start))

	start = time.Now()

	if err = image.Save("renders/test.png"); err != nil {
		panic(err)
	}

	log.Printf("Time passed to save image: %s", time.Since(start))
}
