package standard

import (
	"errors"
	"image/color"
	"math"

	"github.com/GeekchanskiY/avagen/pkg/point"
	"github.com/GeekchanskiY/avagen/pkg/scene"
	"github.com/GeekchanskiY/avagen/pkg/vector"
)

type Generator interface {
	Generate(s *scene.Scene) error
}

type generator struct {
	squareSize int
}

func NewGenerator() Generator {
	return &generator{}
}

func (g *generator) Generate(s *scene.Scene) error {
	if s.Width != s.Height {
		return errors.New("scene width and height must be equal")
	}

	if s.Width < 8 {
		return errors.New("scene width must be greater than 8")
	}

	g.squareSize = s.Width / 8

	triangleColors := g.generateTriangle(g.squareSize, s.FgColor, s.BgColor)
	squareColors := g.generateSquare(g.squareSize, s.FgColor, s.BgColor)
	filledTriangleColors := g.generateFilledTriangle(g.squareSize, s.FgColor, s.BgColor)
	rotatedTriangleColors := g.rotate(triangleColors)
	rotatedRevTriangleColors := g.rotateRev(triangleColors)
	reversedTriangleColors := g.reverse(triangleColors)
	circleColors := g.generateCircle(g.squareSize, s.FgColor, s.BgColor)
	circleFilledColors := g.generateFilledCircle(g.squareSize, s.FgColor, s.BgColor)

	for i := 0; i < g.squareSize; i++ {
		for j := 0; j < g.squareSize; j++ {
			s.Img.Set(i, j, triangleColors[i][j])
			s.Img.Set(i+g.squareSize*1, j, squareColors[i][j])
			s.Img.Set(i+g.squareSize*2, j, filledTriangleColors[i][j])
			s.Img.Set(i+g.squareSize*3, j, rotatedTriangleColors[i][j])
			s.Img.Set(i+g.squareSize*4, j, reversedTriangleColors[i][j])
			s.Img.Set(i+g.squareSize*5, j, rotatedRevTriangleColors[i][j])
			s.Img.Set(i+g.squareSize*6, j, circleColors[i][j])
			s.Img.Set(i+g.squareSize*7, j, circleFilledColors[i][j])
		}
	}

	return nil
}

func (g *generator) generateSquare(squareSize int, fg, _ color.RGBA) [][]color.RGBA {
	squareColors := make([][]color.RGBA, squareSize)
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

func (g *generator) generateFilledTriangle(squareSize int, fg, bg color.RGBA) [][]color.RGBA {
	squareColors := make([][]color.RGBA, squareSize)
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

func (g *generator) generateTriangle(squareSize int, fg, bg color.RGBA) [][]color.RGBA {
	squareColors := make([][]color.RGBA, squareSize)
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

func (g *generator) generateCircle(squareSize int, fg, bg color.RGBA) [][]color.RGBA {
	var (
		radius = (squareSize - 2) / 2
		center = point.Point{
			X: squareSize / 2,
			Y: squareSize / 2,
		}

		circleColors = make([][]color.RGBA, squareSize)
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

func (g *generator) generateFilledCircle(squareSize int, fg, bg color.RGBA) [][]color.RGBA {
	var (
		radius = (squareSize - 2) / 2
		center = point.Point{
			X: squareSize / 2,
			Y: squareSize / 2,
		}

		circleColors = make([][]color.RGBA, squareSize)
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

func (g *generator) rotate(points [][]color.RGBA) [][]color.RGBA {
	var (
		newPoints     [][]color.RGBA
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make([][]color.RGBA, height)
	for i := 0; i < height; i++ {
		newPoints[i] = make([]color.RGBA, width)
		for j := 0; j < width; j++ {
			newPoints[i][j] = points[j][height-i-1]
		}
	}

	return newPoints
}

func (g *generator) rotateRev(points [][]color.RGBA) [][]color.RGBA {
	var (
		newPoints     [][]color.RGBA
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make([][]color.RGBA, height)
	for i := height; i > 0; i-- {
		newPoints[height-i] = make([]color.RGBA, width)
		for j := width; j > 0; j-- {
			newPoints[height-i][width-j] = points[width-j][height-i]
		}
	}

	return newPoints
}

func (g *generator) reverse(points [][]color.RGBA) [][]color.RGBA {
	var (
		newPoints     [][]color.RGBA
		width, height int
	)

	if width = len(points); width == 0 {
		return points
	}

	if height = len(points[0]); height == 0 {
		return points
	}

	newPoints = make([][]color.RGBA, width)

	for i := 0; i < width; i++ {
		newPoints[i] = make([]color.RGBA, height)
		for j := 0; j < height; j++ {
			newPoints[i][j] = points[width-i-1][height-j-1]
		}
	}

	return newPoints
}
