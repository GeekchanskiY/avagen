package standard

import (
	"errors"
	"image/color"
	"math"
	"math/rand"

	"github.com/GeekchanskiY/avagen/pkg/point"
	"github.com/GeekchanskiY/avagen/pkg/scene"
	"github.com/GeekchanskiY/avagen/pkg/vector"
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

type figure [][]color.RGBA

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

func (g *generator) generateRandomFigure() figure {
	switch rand.Intn(8) {
	case 0:
		return g.generateSquare(g.squareSize, g.fgColor, color.RGBA{})
	case 1:
		return g.generateFilledTriangle(g.squareSize, g.fgColor, g.bgColor)
	case 2:
		return g.generateTriangle(g.squareSize, g.fgColor, g.bgColor)
	case 3:
		return g.generateCircle(g.squareSize, g.fgColor, g.bgColor)
	case 4:
		return g.generateFilledCircle(g.squareSize, g.fgColor, g.bgColor)
	default:
		return g.generateSquare(g.squareSize, g.bgColor, color.RGBA{})
	}
}

func (g *generator) generateSquare(squareSize int, fg, _ color.RGBA) figure {
	squareColors := make(figure, squareSize)
	for i := 0; i < squareSize; i++ {
		squareColors[i] = make([]color.RGBA, squareSize)
	}

	for i := range squareColors {
		for j := range squareColors[i] {
			squareColors[i][j] = fg
		}
	}

	return squareColors
}

func (g *generator) generateFilledTriangle(squareSize int, fg, bg color.RGBA) figure {
	squareColors := make(figure, squareSize)
	for i := 0; i < squareSize; i++ {
		squareColors[i] = make([]color.RGBA, squareSize)
	}

	var (
		v0, v1, v2 point.Vertex
	)

	v0 = point.Vertex{
		Position: point.Point{
			X: squareSize / 2,
			Y: 0,
		},
	}

	v1 = point.Vertex{Position: point.Point{X: squareSize, Y: squareSize}}

	v2 = point.Vertex{Position: point.Point{X: 0, Y: squareSize}}

	for y := 0; y < squareSize; y++ {
		for x := 0; x < squareSize; x++ {
			e10 := vector.Edge(v1.Position, v0.Position, point.Point{X: x, Y: y})
			e21 := vector.Edge(v2.Position, v1.Position, point.Point{X: x, Y: y})
			e02 := vector.Edge(v0.Position, v2.Position, point.Point{X: x, Y: y})

			if e10 > 0 && e21 > 0 && e02 > 0 {
				squareColors[x][y] = fg
			} else {
				squareColors[x][y] = bg
			}
		}
	}

	return squareColors
}

func (g *generator) generateTriangle(squareSize int, fg, bg color.RGBA) figure {
	squareColors := make(figure, squareSize)
	for i := 0; i < squareSize; i++ {
		squareColors[i] = make([]color.RGBA, squareSize)
	}

	var (
		v0, v1, v2 point.Vertex
	)

	v0 = point.Vertex{
		Position: point.Point{
			X: squareSize / 2,
			Y: 0,
		},
	}

	v1 = point.Vertex{Position: point.Point{X: squareSize, Y: squareSize}}

	v2 = point.Vertex{Position: point.Point{X: 0, Y: squareSize}}

	for y := 0; y < squareSize; y++ {
		for x := 0; x < squareSize; x++ {
			e10 := vector.Edge(v1.Position, v0.Position, point.Point{X: x, Y: y})
			e21 := vector.Edge(v2.Position, v1.Position, point.Point{X: x, Y: y})
			e02 := vector.Edge(v0.Position, v2.Position, point.Point{X: x, Y: y})

			if e10 >= 0 && e10 <= 12 || e21 == squareSize || e02 >= 0 && e02 <= 12 {
				squareColors[x][y] = fg
			} else {
				squareColors[x][y] = bg
			}
		}
	}

	return squareColors
}

func (g *generator) generateCircle(squareSize int, fg, bg color.RGBA) figure {
	var (
		radius = (squareSize - 2) / 2
		center = point.Point{
			X: squareSize / 2,
			Y: squareSize / 2,
		}

		circleColors = make(figure, squareSize)
	)

	radiusPow := math.Pow(float64(radius), 2)

	for x := 0; x < squareSize; x++ {
		circleColors[x] = make([]color.RGBA, squareSize)
		for y := 0; y < squareSize; y++ {
			diff := math.Pow(float64(center.X-x), 2) + math.Pow(float64(center.Y-y), 2) - radiusPow
			if diff >= float64(-squareSize/2) && diff <= float64(squareSize/2) {
				circleColors[x][y] = fg
			} else {
				circleColors[x][y] = bg
			}
		}
	}

	return circleColors
}

func (g *generator) generateFilledCircle(squareSize int, fg, bg color.RGBA) figure {
	var (
		radius = (squareSize - 2) / 2
		center = point.Point{
			X: squareSize / 2,
			Y: squareSize / 2,
		}

		circleColors = make(figure, squareSize)
	)

	radiusPow := math.Pow(float64(radius), 2)

	for x := 0; x < squareSize; x++ {
		circleColors[x] = make([]color.RGBA, squareSize)
		for y := 0; y < squareSize; y++ {

			if math.Pow(float64(center.X-x), 2)+math.Pow(float64(center.Y-y), 2) <= radiusPow {
				circleColors[x][y] = fg
			} else {
				circleColors[x][y] = bg
			}
		}
	}

	return circleColors
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
