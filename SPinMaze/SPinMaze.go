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

	//The width of the edges.
	lineWidth := 14.0

	//Different maze intialization based on command lines

	// For simulating the simple maze problem in Physarum polycephalum observation, mode="maze";
	// For simulating the shortest path between two cities in China from the road map, mode ="transport"
	mode := os.Args[1]

	//For simulate whether there is light in part of the maze. If there is light. isLight=true. We only add light to the
	//Physarum polycephalum solving maze problem. Else, islight=false.
	isLight := os.Args[2]

	//The fileName is for the name of the output GIF.
	fileName := os.Args[3]
	mu := 1.0
	if mode == "maze" {
		maze = InitializeSimpleMaze(isLight)
		flagD = "noRandom"
	} else if mode == "transport" {
		maze = InitializeTransportMaze()
		flagD = "random"

		//mu is the magnitude of the impact of flux on the conductivity.
		mu = 1.1
	}

	CheckIfIntializeRight(maze)

	//The number of iterations for the maze evolvement simulation.
	numgen := 40
	rand.Seed(time.Now().UnixNano())

	//Start simulation!
	Q := MazeEvolve(maze, numgen, mu, flagD)

	//Draw GIF.
	imageList := DrawMazes(maze, Q, numgen, lineWidth)
	ImagesToGIF(imageList, fileName)

	fmt.Println("Drawing finishes.")
}
