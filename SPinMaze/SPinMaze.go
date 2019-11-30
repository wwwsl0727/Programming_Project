package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	fmt.Println("Maze!")

	var maze Maze
	var flagD string // random, initialize D in random (0,1]
	lineWidth := 14.0
	//Different maze intialization based on command lines.
	mode := os.Args[1]
	isLight := os.Args[2]
	fileName := os.Args[3]
	mu := 1.0
	if mode == "maze" {
		maze = InitializeSimpleMaze(isLight)
		flagD = "noRandom"
	} else if mode == "transport" {
		maze = InitializeTransportMaze()
		flagD = "random"
		mu = 1.1
	}
	CheckIfIntializeRight(maze)
	numgen := 40

	rand.Seed(time.Now().UnixNano())
	//MazeEvolve(maze, numgen, nsteps, t)
	//Q := MazeEvolve(maze, numgen, nsteps, t, flagD)
	Q := MazeEvolve(maze, numgen, mu, flagD)
	imageList := DrawMazes(maze, Q, numgen, lineWidth)
	ImagesToGIF(imageList, fileName)

	fmt.Println("Drawing finishes.")

	/*
		//	Part city
		fmt.Println("maze!")
		var maze Maze
		//Different maze intialization based on command lines.
		var N1, N2, N3, N4, N5, N6, N7, N8 Node

		N1.location.x, N1.location.y = 25.0, 17.5
		N2.location.x, N2.location.y = 40.5, 24.0
		N3.location.x, N3.location.y = 34.5, 19.5
		N4.location.x, N4.location.y = 39.5, 22.0
		N5.location.x, N5.location.y = 38.0, 22.5
		N6.location.x, N6.location.y = 35.0, 24.5
		N7.location.x, N7.location.y = 30.0, 20.0
		N8.location.x, N8.location.y = 25.0, 24.0

		N1.neighbors = append(N1.neighbors, &N7, &N8)
		N2.neighbors = append(N2.neighbors, &N4, &N6)

		N3.neighbors = append(N3.neighbors, &N4, &N5, &N6, &N7)
		N4.neighbors = append(N4.neighbors, &N2, &N3, &N5)
		N5.neighbors = append(N5.neighbors, &N3, &N4, &N6, &N7)
		N6.neighbors = append(N6.neighbors, &N3, &N5, &N2, &N7, &N8)
		N7.neighbors = append(N7.neighbors, &N3, &N5, &N6, &N1, &N8)
		N8.neighbors = append(N8.neighbors, &N6, &N1, &N7)

		N1.name, N2.name, N3.name, N4.name, N5.name, N6.name = "N1", "N2", "N3", "N4", "N5", "N6"
		N7.name, N8.name = "N7", "N8"

		maze = append(maze, &N1, &N2, &N3, &N4, &N5, &N6, &N7, &N8)
		//fmt.Println(len(transportMaze))
		CheckIfIntializeRight(maze)
		numgen := 50
		//nsteps := 1 // This is for calculating conductivity

		//t := 1.0
		lineWidth := 14.0
		fileName := "partcity_test"
		//Q := MazeEvolve(maze, numgen, nsteps, t, "random")
		Q := MazeEvolve(maze, numgen, 1.0, "random")
		imageList := DrawMazes(maze, Q, numgen, lineWidth)
		ImagesToGIF(imageList, fileName)
		fmt.Println("finish!")
	*/
}
