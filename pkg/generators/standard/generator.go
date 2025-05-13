package standard

import (
	"errors"
	"fmt"
	"image/color"

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

	for i := 0; i < g.squareSize; i++ {
		for j := 0; j < g.squareSize; j++ {
			s.Img.Set(i, j, triangleColors[i][j])
			s.Img.Set(i+g.squareSize*1, j, squareColors[i][j])
			s.Img.Set(i+g.squareSize*2, j, filledTriangleColors[i][j])
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

			fmt.Println(e10, e21, e02)

			if e10 >= 0 && e10 <= 12 || e21 == 25 || e02 >= 0 && e02 <= 12 {
				squareColors[x][y] = fg
			} else {
				squareColors[x][y] = bg
			}
		}
	}

	return squareColors
}
