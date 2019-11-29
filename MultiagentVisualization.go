package main

import (
	"image"
	"image/png"
	"os"

	"code.google.com/p/draw2d/draw2d"
)

type drawer struct {
	gc     *draw2d.ImageGraphicContext
	img    image.Image
	width  int
	height int
}

//
// func DrawGameBoards(boards []Board, cellWidth int) []image.Image {
// 	numGenerations := len(boards)
// 	imageList := make([]image.Image, numGenerations)
// 	for i := range boards {
// 		imageList[i] = DrawGameBoard(boards[i], cellWidth)
// 	}
// 	return imageList
// }

func OutPNGPictures(boards []multiAgentMatrix, cellWidth int, CN float64) {
	imglist := DrawGameBoards(boards, 5, CN)
	// Images to severial PNG(imglist, "outputFile")
	for i := range imglist {
		f, err := os.Create("i" + "Prisoners.png")
		if err != nil {
			panic(err)
		}
		png.Encode(f, imglist[i])
	}
}

func DrawGameBoards(boards []multiAgentMatrix, cellWidth int, CN float64) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth, CN) //CN is here to check if it's good food or bad food.
	}
	return imageList
}

func DrawGameBoard(board multiAgentMatrix, cellWidth int, CN float64) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth

	// declare colors
	// black := MakeColor(0, 0, 0)
	grey := MakeColor(211, 211, 211)  //background
	red := MakeColor(0, 0, 0)         //bad food
	green := MakeColor(0, 255, 255)   //good food
	yellow := MakeColor(255, 255, 0)  //Agent
	white := MakeColor(255, 255, 255) //light

	c := CreateNewMazes(width, height) //Create new drawer for board
	//set the entire board as grey
	c.SetFillColor(grey)
	// c.ClearRect(0, 0, height, width)
	c.Clear()

	// draw the grid lines in white
	// c.SetStrokeColor(white)
	// DrawGridLines(c, cellWidth)

	// fill in colored squares and draw
	for i := range board {
		for j := range board[i] {
			if board[i][j].IsAgent {
				c.SetFillColor(yellow)
				c.Circle(float64(i), float64(j), 1.0) //draw a circle of radius 1
				c.Fill()

			} else if board[i][j].IsFood {
				if board[i][j].foodChemo == CN {
					c.SetFillColor(green)                         //good food
					c.Rectangle(float64(i), float64(j), 1.0, 1.0) //draw a square of length 1
					c.Fill()
				} else if board[i][j].foodChemo == 0 {
					c.SetFillColor(red) // bad food
					c.Rectangle(float64(i), float64(j), 1.0, 1.0)
					c.Fill()
				}
			}
			if board[i][j].haslight {
				c.SetFillColor(white)
				c.Rectangle(float64(i), float64(j), 1, 1)
				c.Fill()
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}
	return c.img
}

// func DrawGridLines(pic Canvas, cellWidth int) {
// 	w, h := pic.width, pic.height
// 	// first, draw vertical lines
// 	for i := 1; i < pic.width/cellWidth; i++ {
// 		y := i * cellWidth
// 		pic.MoveTo(0.0, float64(y))
// 		pic.LineTo(float64(w), float64(y))
// 	}
// 	// next, draw horizontal lines
// 	for j := 1; j < pic.height/cellWidth; j++ {
// 		x := j * cellWidth
// 		pic.MoveTo(float64(x), 0.0)
// 		pic.LineTo(float64(x), float64(h))
// 	}
// 	pic.Stroke()
// }
