package main

import (
	"image"
)

//
// func DrawGameBoards(boards []Board, cellWidth int) []image.Image {
// 	numGenerations := len(boards)
// 	imageList := make([]image.Image, numGenerations)
// 	for i := range boards {
// 		imageList[i] = DrawGameBoard(boards[i], cellWidth)
// 	}
// 	return imageList
// }

func DrawGameBoard(board multiAgentMatrix, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewCanvas(width, height)

	// declare colors
	// black := MakeColor(0, 0, 0)
	black := MakeColor(0, 0, 0)
	yellow := MakeColor(255, 255, 0)
	white := MakeColor(255, 255, 255)
	/*
		//set the entire board as black
		c.SetFillColor(gray)
		c.ClearRect(0, 0, height, width)
		c.Clear()
	*/

	// draw the grid lines in white
	c.SetStrokeColor(white)
	DrawGridLines(c, cellWidth)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].IsFood {
				c.SetFillColor(black)
			} else if board[i][j].IsAgent {
				c.SetFillColor(yellow)
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}
	return c.img
}

func DrawGridLines(pic Canvas, cellWidth int) {
	w, h := pic.width, pic.height
	// first, draw vertical lines
	for i := 1; i < pic.width/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < pic.height/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}
