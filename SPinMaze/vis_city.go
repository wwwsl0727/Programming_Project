package main

import (
	"fmt"
	"gogif"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"

	"code.google.com/p/draw2d/draw2d"
)

type Mazes struct {
	gc     *draw2d.ImageGraphicContext
	img    image.Image
	width  int
	height int
}

func MakeColor(r, g, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}

func (m *Mazes) MoveTo(x, y float64) {
	m.gc.MoveTo(x, y)
}

func (m *Mazes) LineTo(x, y float64) {
	m.gc.LineTo(x, y)
}

func (m *Mazes) DrawLine(source OrderedPair, destination OrderedPair, color color.Color, lineWidth float64) {
	m.SetStrokeColor(color)
	m.SetLineWidth(lineWidth)
	m.MoveTo(source.x, source.y)
	m.LineTo(destination.x, destination.y)
	m.FillStroke()
}

func (m *Mazes) SetStrokeColor(col color.Color) {
	m.gc.SetStrokeColor(col)
}

func (m *Mazes) SetFillColor(col color.Color) {
	m.gc.SetFillColor(col)
}

func (m *Mazes) SetLineWidth(w float64) {
	m.gc.SetLineWidth(w)
}

func (m *Mazes) Stroke() {
	m.gc.Stroke()
}

func (m *Mazes) Fill() {
	m.gc.Fill()
}

func (m *Mazes) FillStroke() {
	m.gc.FillStroke()
}

func (m *Mazes) BeginPath() {
	m.gc.BeginPath()

}

func (m *Mazes) Circle(cx, cy, r float64) {
	m.gc.ArcTo(cx, cy, r, r, 0, -math.Pi*2)
	m.gc.Close()
}

func (m *Mazes) Rectangle(x, y, w, h float64) {
	x = x - w/2
	y = y - h/2
	m.gc.MoveTo(x, y)
	m.gc.LineTo(x+w, y)
	m.gc.LineTo(x+w, y+h)
	m.gc.LineTo(x, y+h)
	m.gc.LineTo(x, y)
}

func CreateNewMazes(w, h int) Mazes {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background
	gc.Clear()
	gc.SetFillColor(image.Black)

	return Mazes{gc, i, w, h}
}

func MazeScaling(maze Maze, scale float64) Maze {
	for i := 0; i < len(maze); i++ {
		maze[i].location.x = scale*maze[i].location.x + 30
		maze[i].location.y = scale*maze[i].location.y + 30
	}
	return maze
}

func DrawMazes(maze Maze, flow []Matrix, numGen int, lineWidth float64) []image.Image {
	imageList := make([]image.Image, numGen)
	maze = MazeScaling(maze, 30)
	for i := 0; i < numGen; i++ {
		imageList[i] = DrawMaze(maze, flow[i], i, lineWidth)
	}
	return imageList
}

func IsNeighbor(maze Maze, a, b int) bool {
	for _, neighbor := range maze[a].neighbors {
		if neighbor == maze[b] {
			return true
		}
	}
	return false
}

func DrawMaze(maze Maze, currFlow Matrix, numGen int, lineWidth float64) image.Image {
	// yellow := MakeColor(255, 255, 0)
	red := MakeColor(255, 0, 0)
	blue := MakeColor(0, 0, 255)
	cyan := MakeColor(0, 255, 255)
	black := MakeColor(0, 0, 0)
	pathWidth := lineWidth / 3.0

	maximumX := -214748364748.0
	maximumY := -214748364748.0
	for i := 0; i < len(maze); i++ {
		if maze[i].location.x > maximumX {
			maximumX = maze[i].location.x
		}
		if maze[i].location.y > maximumY {
			maximumY = maze[i].location.y
		}
	}
	width := int(1.5 * maximumX)
	height := int(1.5 * maximumY)
	nodeW := 20.0
	nodeH := 20.0
	r := 15.0
	m := CreateNewMazes(width, height)

	for _, node := range maze {
		for _, neighbor := range node.neighbors {
			if neighbor != nil {
				m.SetStrokeColor(black)
				m.SetLineWidth(lineWidth)
				m.MoveTo(node.location.x, node.location.y)
				m.LineTo(neighbor.location.x, neighbor.location.y)
				m.FillStroke()
			}
		}
	}
	for i := range maze {
		if maze[i].name == "N1" {
			m.SetFillColor(red)
			m.Circle(maze[i].location.x, maze[i].location.y, r)
			m.Fill()
		} else if maze[i].name == "N2" {
			m.SetFillColor(blue)
			m.Circle(maze[i].location.x, maze[i].location.y, r)
			m.Fill()
		} else {
			m.SetFillColor(cyan)
			m.Rectangle(maze[i].location.x, maze[i].location.y, nodeW, nodeH)
			m.Fill()
		}
	}
	if numGen == 0 {
		return m.img
	}
	for i := 0; i < len(currFlow); i++ {
		for j := i; j < len(currFlow[i]); j++ {

			if math.Abs(math.Abs(currFlow[i][j])-1.0) < 0.9 {
				if IsNeighbor(maze, i, j) {
					m.DrawLine(maze[i].location, maze[j].location, red, pathWidth)
				}
			} else if math.Abs(currFlow[i][j]) < 1.0 {
				color := MakeColor(150, 200, 0)
				if IsNeighbor(maze, i, j) {
					m.DrawLine(maze[i].location, maze[j].location, color, pathWidth)
				}

			} else if math.Abs(currFlow[i][j]) < 0.0000001 {
				if IsNeighbor(maze, i, j) {
					m.DrawLine(maze[i].location, maze[j].location, black, pathWidth)
				}
			}

		}
	}
	for i := range maze {
		if maze[i].name == "N1" {
			m.SetFillColor(red)
			m.Circle(maze[i].location.x, maze[i].location.y, r)
			m.Fill()
		} else if maze[i].name == "N2" {
			m.SetFillColor(blue)
			m.Circle(maze[i].location.x, maze[i].location.y, r)
			m.Fill()
		} else {
			m.SetFillColor(cyan)
			m.Rectangle(maze[i].location.x, maze[i].location.y, nodeW, nodeH)
			m.Fill()
		}
	}

	return m.img
}

func ImagesToGIF(imglist []image.Image, filename string) {

	// get ready to write images to files

	w, err := os.Create(filename + ".out.gif")

	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}

	defer w.Close()
	var g gif.GIF
	g.Delay = make([]int, len(imglist))
	g.Image = make([]*image.Paletted, len(imglist))
	g.LoopCount = 10

	for i := range imglist {

		g.Image[i] = ImageToPaletted(imglist[i])
		g.Delay[i] = 1

	}

	gif.EncodeAll(w, &g)
}
func ImageToPaletted(img image.Image) *image.Paletted {

	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()

		pm = image.NewPaletted(b, nil)

		q := &gogif.MedianCutQuantizer{NumColor: 256}

		q.Quantize(pm, b, img, image.ZP)

	}

	return pm
}
