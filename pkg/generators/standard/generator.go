package standard

import (
	"errors"
	"image/color"
	"math/rand"

	"github.com/GeekchanskiY/avagen/pkg/scene"
)

type Generator interface {
	Generate(s *scene.Scene, numSquares int) error
}

type generator struct {
	squareSize       int
	bgColor, fgColor color.RGBA
}

func NewGenerator() Generator {
	return &generator{}
}

// Generate takes a scene which it will edit, and numSquares in row
func (g *generator) Generate(s *scene.Scene, numSquares int) error {
	if s.Width != s.Height {
		return errors.New("scene width and height must be equal")
	}

	if s.Width < 8 {
		return errors.New("scene width must be greater than 8")
	}

	g.squareSize = s.Width / numSquares

	g.bgColor = getRandomBgColor()
	g.fgColor = getRandomFgColor()

	s.BgColor = g.bgColor
	s.FgColor = g.fgColor

	s.Clear()

	// Generate figures
	figures := make([][]figure, numSquares)
	for i := 0; i < numSquares; i++ {
		figures[i] = make([]figure, numSquares)
		for j := 0; j < numSquares; j++ {
			figures[i][j] = g.generateRandomFigure()
			switch rand.Intn(5) {
			case 2:
				figures[i][j] = g.reverse(figures[i][j])
			case 1:
				figures[i][j] = g.rotate(figures[i][j])
			case 0:
				figures[i][j] = g.rotateRev(figures[i][j])
			default:
			}
		}

	}

	// Draw figures on the scene
	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			for x := 0; x < g.squareSize; x++ {
				for y := 0; y < g.squareSize; y++ {
					s.Img.Set(x+i*g.squareSize, y+j*g.squareSize, figures[i][j][x][y])
				}
			}
		}
	}

	return nil
}

func (g *generator) rotate(points figure) figure {
	var (
		newPoints     figure
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make(figure, height)
	for i := 0; i < height; i++ {
		newPoints[i] = make([]color.RGBA, width)
		for j := 0; j < width; j++ {
			newPoints[i][j] = points[j][height-i-1]
		}
	}

	return newPoints
}

func (g *generator) rotateRev(points figure) figure {
	var (
		newPoints     figure
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make(figure, height)
	for i := height; i > 0; i-- {
		newPoints[height-i] = make([]color.RGBA, width)
		for j := width; j > 0; j-- {
			newPoints[height-i][width-j] = points[width-j][height-i]
		}
	}

	return newPoints
}

func (g *generator) reverse(points figure) figure {
	var (
		newPoints     figure
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make(figure, width)

	for i := 0; i < width; i++ {
		newPoints[i] = make([]color.RGBA, height)
		for j := 0; j < height; j++ {
			newPoints[i][j] = points[width-i-1][height-j-1]
		}
	}

	return newPoints
}
