package standard

import (
	"image/color"
	"math/rand"
)

// Default colors are taken from the official Gruvbox colorscheme
var (
	colorBg1 = color.RGBA{
		R: 29,
		G: 32,
		B: 33,
		A: 255,
	}
	colorBg2 = color.RGBA{
		R: 40,
		G: 40,
		B: 40,
		A: 255,
	}
	colorBg3 = color.RGBA{
		R: 50,
		G: 48,
		B: 47,
		A: 255,
	}
	// colorFg1 - red
	colorFg1 = color.RGBA{
		R: 204,
		G: 36,
		B: 29,
		A: 255,
	}
	// colorFg2 - yellow
	colorFg2 = color.RGBA{
		R: 215,
		G: 153,
		B: 33,
		A: 255,
	}
	// colorFg3 - blue
	colorFg3 = color.RGBA{
		R: 69,
		G: 133,
		B: 136,
		A: 255,
	}
	// colorFg4 - green (aqua)
	colorFg4 = color.RGBA{
		R: 104,
		G: 157,
		B: 106,
		A: 255,
	}
)

func getRandomFgColor() color.RGBA {
	colors := []color.RGBA{colorFg1, colorFg2, colorFg3, colorFg4}

	return colors[rand.Intn(len(colors))]
}

func getRandomBgColor() color.RGBA {
	colors := []color.RGBA{colorBg1, colorBg2, colorBg3}

	return colors[rand.Intn(len(colors))]
}
