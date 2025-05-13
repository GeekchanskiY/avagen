package standard

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/GeekchanskiY/avagen/pkg/point"
	"github.com/GeekchanskiY/avagen/pkg/vector"
)

type figure [][]color.RGBA

// generateRandomFigure generates a random figure(triangle, circle, square, nil)
func (g *generator) generateRandomFigure() figure {
	switch rand.Intn(8) {
	case 0:
		return g.generateSquare(g.squareSize, g.fgColor, color.RGBA{})
	case 1:
		return g.generateFilledTriangle(g.squareSize, g.fgColor, g.bgColor)
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
