package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Maze []*Node

type Node struct {
	neighbors []*Node
	pressure  float64
	location  OrderedPair
	name      string
	city      string
}

type OrderedPair struct {
	x, y float64
}

type Matrix [][]float64

func main() {
	fmt.Println("Maze!")

	var maze Maze
	var lineWidth float64
	//Different maze intialization based on command lines.
	mode := os.Args[1]
	isLight := os.Args[2]
	fileName := os.Args[3]
	if mode == "maze" {
		maze = InitializeSimpleMaze(isLight)
		lineWidth = 14.0
	} else if mode == "transport" {
		maze = InitializeTransportMaze()
		lineWidth = 9.0
	}
	CheckIfIntializeRight(maze)
	numgen := 30

	nsteps := 10000 // This is for calculating conductivity

	t := 0.0001

	rand.Seed(time.Now().UnixNano())
	//MazeEvolve(maze, numgen, nsteps, t)
	Q := MazeEvolve(maze, numgen, nsteps, t)

	imageList := DrawMazes(maze, Q, numgen, lineWidth)
	ImagesToGIF(imageList, fileName)

	fmt.Println("Drawing finishes.")
}

func CheckIfIntializeRight(maze Maze) {
	//Check if neighbors of maze match.
	//For each node
	for i := range maze {
		//address of neighbors of Ni
		currNode := *maze[i]
		neighbors := currNode.neighbors
		//Check if Ni is also the neighbor of neighbors[j]
		for j := range neighbors {
			neighborOfNei := (*neighbors[j]).neighbors
			flag := 1 // currNode is not the neighbor of node j
			for k := range neighborOfNei {
				neighborOfNeiName := neighborOfNei[k].name
				if currNode.name == neighborOfNeiName {
					flag = 0
				}
			}
			if flag == 1 {
				fmt.Printf("Initialization wrong, the neighbor of %s should contain %s!\n", (*neighbors[j]).name, currNode.name)
				os.Exit(1)
			}
			//neighbornames = append(neighbornames, (*neighbors[j]).name)
		}

	}
	/*
		fmt.Println("address of N1", maze[0])                            // &{[0xc000086180] 0 {0 0} N1}
		fmt.Println("attributes of N1", *maze[0])                        //{[0xc000086180] 0 {0 0} N1}
		fmt.Println("name of N1", (*maze[0]).name)                       //"N1"
		fmt.Println(" address of neighbors of N1", (*maze[0]).neighbors) //[&N6] [0xc0000980f0]
		address := (*maze[0]).neighbors[0]                               //&N6
		fmt.Println("address of N6", address)                            //&N6
		fmt.Println("name of N6", (*address).name)                       //"N6"

		Check if the intialization is right
		For each node
		for i := range maze {
			//address of neighbors of Ni
			neighbors := (*maze[i]).neighbors
			var neighbornames []string
			for j := range neighbors {
				neighbornames = append(neighbornames, (*neighbors[j]).name)
			}
			fmt.Println("location and neighbor names of each node", (*maze[i]).name, (*maze[i]).location, neighbornames)

		}
	*/
}

func InitializeTransportMaze() Maze {
	var transportMaze Maze
	var N1, N2, N3, N4, N5, N6, N7, N8, N9, N10, N11, N12, N13, N14, N15, N16, N17, N18, N19 Node
	var N20, N21, N22, N23, N24, N25, N26, N27, N28, N29, N30, N31 Node
	// N1.location.x, N1.location.y = 3700.0, 1350.0
	// N2.location.x, N2.location.y = 3750.0, 1400.0
	// N3.location.x, N3.location.y = 4200.0, 2300.0
	// N4.location.x, N4.location.y = 2750.0, 2500.0
	// N5.location.x, N5.location.y = 3200.0, 1250.0
	// N6.location.x, N6.location.y = 1250.0, 700.0
	// N7.location.x, N7.location.y = 1200.0, 2350.0
	// N8.location.x, N8.location.y = 2800.0, 1500.0
	// N9.location.x, N9.location.y = 2950.0, 3250.0
	// N10.location.x, N10.location.y = 4400.0, 550.0
	// N11.location.x, N11.location.y = 4300.0, 800.0
	// N12.location.x, N12.location.y = 4200.0, 1000.0
	// N13.location.x, N13.location.y = 3500.0, 1500.0
	// N14.location.x, N14.location.y = 3300.0, 1550.0
	// N15.location.x, N15.location.y = 2400.0, 1750.0
	// N16.location.x, N16.location.y = 3750.0, 1700.0
	// N17.location.x, N17.location.y = 3450.0, 1950.0
	// N18.location.x, N18.location.y = 3950.0, 2200.0
	// N19.location.x, N19.location.y = 3800.0, 2250.0
	// N20.location.x, N20.location.y = 4050.0, 2400.0
	// N21.location.x, N21.location.y = 4000.0, 2800.0
	// N22.location.x, N22.location.y = 3750.0, 2600.0
	// N23.location.x, N23.location.y = 3400.0, 2750.0
	// N24.location.x, N24.location.y = 3500.0, 2450.0
	// N25.location.x, N25.location.y = 3450.0, 3250.0
	// N26.location.x, N26.location.y = 3150.0, 3550.0
	// N27.location.x, N27.location.y = 2500.0, 1750.0
	// N28.location.x, N28.location.y = 3000.0, 2000.0
	// N29.location.x, N29.location.y = 2500.0, 2400.0
	// N30.location.x, N30.location.y = 2800.0, 2900.0
	// N31.location.x, N31.location.y = 2400.0, 3000.0
	//N1 becomes N27 lanzhou, N2 becomes N20 Hangzhou.
	//
	//N1.location.x, N1.location.y = 37.0, 13.5
	//N2.location.x, N2.location.y = 37.5, 14.0
	N3.location.x, N3.location.y = 42.0, 23.0
	N4.location.x, N4.location.y = 27.5, 25.0
	N5.location.x, N5.location.y = 32.0, 12.5
	N6.location.x, N6.location.y = 12.5, 7.0
	N7.location.x, N7.location.y = 12.0, 23.5
	N8.location.x, N8.location.y = 28.0, 15.0
	N9.location.x, N9.location.y = 29.5, 32.5
	N10.location.x, N10.location.y = 44.0, 5.5
	N11.location.x, N11.location.y = 43.0, 8.0
	N12.location.x, N12.location.y = 42.0, 10.0
	N13.location.x, N13.location.y = 35.0, 15.0
	N14.location.x, N14.location.y = 33.0, 15.5
	N15.location.x, N15.location.y = 24.0, 17.5
	N16.location.x, N16.location.y = 37.5, 17.0
	N17.location.x, N17.location.y = 34.5, 19.5
	N18.location.x, N18.location.y = 39.5, 22.0
	N19.location.x, N19.location.y = 38.0, 22.5
	//N20.location.x, N20.location.y = 40.5, 24.0
	N21.location.x, N21.location.y = 40.0, 28.0
	N22.location.x, N22.location.y = 37.5, 26.0
	N23.location.x, N23.location.y = 34.0, 27.5
	N24.location.x, N24.location.y = 35.0, 24.5
	N25.location.x, N25.location.y = 34.5, 32.5
	N26.location.x, N26.location.y = 31.5, 35.5
	//N27.location.x, N27.location.y = 25.0, 17.5
	N28.location.x, N28.location.y = 30.0, 20.0
	N29.location.x, N29.location.y = 25.0, 24.0
	N30.location.x, N30.location.y = 28.0, 29.0
	N31.location.x, N31.location.y = 24.0, 30.0

	N1.location.x, N1.location.y = 25.0, 17.5
	N2.location.x, N2.location.y = 40.5, 24.0
	N20.location.x, N20.location.y = 37.5, 14.0
	N27.location.x, N27.location.y = 37.0, 13.5

	//N1.neighbors = append(N1.neighbors, &N2, &N5, &N12, &N13)
	//N2.neighbors = append(N2.neighbors, &N1, &N3, &N12, &N13, &N16, &N18)
	N3.neighbors = append(N3.neighbors, &N2, &N16, &N18, &N20, &N21)
	//N4.neighbors = append(N4.neighbors, &N14, &N16, &N17, &N19, &N20, &N22, &N23, &N24, &N27, &N28, &N29, &N30)
	//N4.neighbors = append(N4.neighbors, &N14, &N16, &N17, &N19, &N2, &N22, &N23, &N24, &N27, &N28, &N29, &N30)
	N4.neighbors = append(N4.neighbors, &N14, &N16, &N17, &N19, &N2, &N22, &N23, &N24, &N1, &N28, &N29, &N30)
	//N5.neighbors = append(N5.neighbors, &N1, &N8, &N14, &N28)
	N5.neighbors = append(N5.neighbors, &N27, &N8, &N14, &N28)
	//N6.neighbors = append(N6.neighbors, &N7, &N8, &N14, &N15, &N27)
	N6.neighbors = append(N6.neighbors, &N7, &N8, &N14, &N15, &N1)
	N7.neighbors = append(N7.neighbors, &N6, &N15, &N29, &N31)
	//N8.neighbors = append(N8.neighbors, &N5, &N6, &N14, &N27, &N28)
	N8.neighbors = append(N8.neighbors, &N5, &N6, &N14, &N1, &N28)
	//N9.neighbors = append(N9.neighbors, &N20, &N22, &N23, &N24, &N25, &N26, &N30, &N31)
	N9.neighbors = append(N9.neighbors, &N2, &N22, &N23, &N24, &N25, &N26, &N30, &N31)
	N10.neighbors = append(N10.neighbors, &N11)
	N11.neighbors = append(N11.neighbors, &N10, &N12)
	//N12.neighbors = append(N12.neighbors, &N1, &N2, &N11, &N16, &N18)
	N12.neighbors = append(N12.neighbors, &N27, &N20, &N11, &N16, &N18)
	//N13.neighbors = append(N13.neighbors, &N1, &N2, &N14, &N16, &N17, &N19)
	N13.neighbors = append(N13.neighbors, &N27, &N20, &N14, &N16, &N17, &N19)
	// N14.neighbors = append(N14.neighbors, &N4, &N5, &N6, &N8, &N13, &N17, &N23, &N27, &N28)
	// N15.neighbors = append(N15.neighbors, &N6, &N7, &N27)
	N14.neighbors = append(N14.neighbors, &N4, &N5, &N6, &N8, &N13, &N17, &N23, &N1, &N28)
	N15.neighbors = append(N15.neighbors, &N6, &N7, &N1)
	//N16.neighbors = append(N16.neighbors, &N2, &N3, &N4, &N12, &N13, &N17, &N18, &N19, &N24)
	N16.neighbors = append(N16.neighbors, &N20, &N3, &N4, &N12, &N13, &N17, &N18, &N19, &N24)
	N17.neighbors = append(N17.neighbors, &N4, &N13, &N14, &N16, &N18, &N19, &N22, &N23, &N24, &N28)
	N18.neighbors = append(N18.neighbors, &N2, &N3, &N12, &N16, &N17, &N19, &N20)
	N19.neighbors = append(N19.neighbors, &N4, &N13, &N16, &N17, &N18, &N21, &N22, &N24, &N28)
	//N20.neighbors = append(N20.neighbors, &N3, &N4, &N9, &N18, &N21, &N22, &N23, &N24, &N25, &N30)
	//N21.neighbors = append(N21.neighbors, &N3, &N19, &N20, &N22, &N23, &N24, &N25, &N30)
	N21.neighbors = append(N21.neighbors, &N3, &N19, &N2, &N22, &N23, &N24, &N25, &N30)
	// N22.neighbors = append(N22.neighbors, &N4, &N9, &N17, &N19, &N20, &N21, &N23, &N24, &N25, &N28, &N30)
	// N23.neighbors = append(N23.neighbors, &N4, &N9, &N14, &N17, &N20, &N21, &N22, &N24, &N25, &N28, &N30)
	N22.neighbors = append(N22.neighbors, &N4, &N9, &N17, &N19, &N2, &N21, &N23, &N24, &N25, &N28, &N30)
	N23.neighbors = append(N23.neighbors, &N4, &N9, &N14, &N17, &N2, &N21, &N22, &N24, &N25, &N28, &N30)
	// N24.neighbors = append(N24.neighbors, &N4, &N9, &N16, &N17, &N19, &N20, &N21, &N22, &N23, &N25, &N28, &N29, &N30)
	// N25.neighbors = append(N25.neighbors, &N9, &N20, &N21, &N22, &N23, &N24, &N26, &N30)
	N24.neighbors = append(N24.neighbors, &N4, &N9, &N16, &N17, &N19, &N2, &N21, &N22, &N23, &N25, &N28, &N29, &N30)
	N25.neighbors = append(N25.neighbors, &N9, &N2, &N21, &N22, &N23, &N24, &N26, &N30)
	N26.neighbors = append(N26.neighbors, &N9, &N25)
	//N27.neighbors = append(N27.neighbors, &N4, &N6, &N8, &N14, &N15, &N28, &N29)
	//N28.neighbors = append(N28.neighbors, &N4, &N5, &N8, &N14, &N17, &N19, &N22, &N23, &N24, &N27, &N29, &N31)
	N28.neighbors = append(N28.neighbors, &N4, &N5, &N8, &N14, &N17, &N19, &N22, &N23, &N24, &N1, &N29, &N31)
	// N29.neighbors = append(N29.neighbors, &N4, &N7, &N24, &N27, &N28, &N31)
	N29.neighbors = append(N29.neighbors, &N4, &N7, &N24, &N1, &N28, &N31)
	// N30.neighbors = append(N30.neighbors, &N4, &N9, &N20, &N21, &N22, &N23, &N24, &N25, &N31)
	N30.neighbors = append(N30.neighbors, &N4, &N9, &N2, &N21, &N22, &N23, &N24, &N25, &N31)
	N31.neighbors = append(N31.neighbors, &N7, &N9, &N28, &N29, &N30)

	N1.neighbors = append(N1.neighbors, &N4, &N6, &N8, &N14, &N15, &N28, &N29)
	N2.neighbors = append(N2.neighbors, &N3, &N4, &N9, &N18, &N21, &N22, &N23, &N24, &N25, &N30)
	N20.neighbors = append(N20.neighbors, &N27, &N3, &N12, &N13, &N16, &N18)
	N27.neighbors = append(N27.neighbors, &N20, &N5, &N12, &N13)

	// N1.city, N2.city, N3.city, N4.city, N5.city, N6.city = "Beijing", "Tianjin", "Shanghai", "Chongqing", "Huhhot", "Umumqi"
	N1.city, N2.city, N3.city, N4.city, N5.city, N6.city = "Lanzhou", "Hangzhou", "Shanghai", "Chongqing", "Huhhot", "Umumqi"
	N7.city, N8.city, N9.city, N10.city, N11.city, N12.city = "Lhasa", "Yinchuan", "Nannin", "Harbin", "Changchun", "Shenyang"
	N13.city, N14.city, N15.city, N16.city, N17.city, N18.city, N19.city = "Shijiazhuang", "Taiyuan", "Xining", "Jinan", "Zhengzhou", "Nanjing", "Hefei"
	// N20.city, N21.city, N22.city, N23.city, N24.city, N25.city, N26.city = "Hangzhou", "Fuzhou", "Nanchang", "Changsha", "Wuhan", "Guangzhou", "Haikou"
	// N27.city, N28.city, N29.city, N30.city, N31.city = "Lanzhou", "Sian", "Chengdu", "Guiyang", "Kunming"
	N20.city, N21.city, N22.city, N23.city, N24.city, N25.city, N26.city = "Tianjin", "Fuzhou", "Nanchang", "Changsha", "Wuhan", "Guangzhou", "Haikou"
	N27.city, N28.city, N29.city, N30.city, N31.city = "Beijing", "Sian", "Chengdu", "Guiyang", "Kunming"

	N1.name, N2.name, N3.name, N4.name, N5.name, N6.name = "N1", "N2", "N3", "N4", "N5", "N6"
	N7.name, N8.name, N9.name, N10.name, N11.name, N12.name = "N7", "N8", "N9", "N10", "N11", "N12"
	N13.name, N14.name, N15.name, N16.name, N17.name, N18.name, N19.name = "N13", "N14", "N15", "N16", "N17", "N18", "N19"
	N20.name, N21.name, N22.name, N23.name, N24.name, N25.name, N26.name = "N20", "N21", "N22", "N23", "N24", "N25", "N26"
	N27.name, N28.name, N29.name, N30.name, N31.name = "N27", "N28", "N29", "N30", "N31"

	transportMaze = append(transportMaze, &N1, &N2, &N3, &N4, &N5, &N6, &N7, &N8, &N9, &N10, &N11, &N12, &N13, &N14, &N15, &N16, &N17, &N18, &N19, &N20,
		&N21, &N22, &N23, &N24, &N25, &N26, &N27, &N28, &N29, &N30, &N31)
	//fmt.Println(len(transportMaze))
	return transportMaze
}

func InitializeSimpleMaze(isLight string) Maze {
	//Initialize nodes
	var maze Maze
	if isLight == "false" {

		var N1, N2, N3, N4, N5, N6, N7, N8, N9, N10, N11, N12, N13, N14, N15, N16, N17, N18, N19 Node

		N1.location.x, N1.location.y = 0.5, 13.0
		N2.location.x, N2.location.y = 10.5, 13.0
		N3.location.x, N3.location.y = 13.5, 11.0
		N4.location.x, N4.location.y = 11.5, 0.0
		N5.location.x, N5.location.y = 9.5, 0.0
		N6.location.x, N6.location.y = 0.5, 6.0
		N7.location.x, N7.location.y = 0.5, 1.0
		N8.location.x, N8.location.y = 0.0, 1.0
		N9.location.x, N9.location.y = 0.0, 0.0
		N10.location.x, N10.location.y = 1.5, 6.0
		N11.location.x, N11.location.y = 1.5, 2.0
		N12.location.x, N12.location.y = 9.5, 2.0
		N13.location.x, N13.location.y = 16.5, 0.0
		N14.location.x, N14.location.y = 16.5, 9.0
		N15.location.x, N15.location.y = 11.5, 10.0
		N16.location.x, N16.location.y = 13.5, 9.0
		N17.location.x, N17.location.y = 9.0, 10.0
		N18.location.x, N18.location.y = 9.0, 11.0
		N19.location.x, N19.location.y = 13.5, 13.0

		N1.neighbors = append(N1.neighbors, &N6)
		N2.neighbors = append(N2.neighbors, &N19)
		N3.neighbors = append(N3.neighbors, &N16, &N18, &N19)
		N4.neighbors = append(N4.neighbors, &N5, &N13, &N15)
		N5.neighbors = append(N5.neighbors, &N4, &N9, &N12)
		N6.neighbors = append(N6.neighbors, &N1, &N10, &N7)
		N7.neighbors = append(N7.neighbors, &N6, &N8)
		N8.neighbors = append(N8.neighbors, &N7, &N9)
		N9.neighbors = append(N9.neighbors, &N5, &N8)
		N10.neighbors = append(N10.neighbors, &N6, &N11)
		N11.neighbors = append(N11.neighbors, &N10, &N12)
		N12.neighbors = append(N12.neighbors, &N5, &N11)
		N13.neighbors = append(N13.neighbors, &N4, &N14)
		N14.neighbors = append(N14.neighbors, &N13, &N16)
		N15.neighbors = append(N15.neighbors, &N4, &N17)
		N16.neighbors = append(N16.neighbors, &N3, &N14)
		N17.neighbors = append(N17.neighbors, &N15, &N18)
		N18.neighbors = append(N18.neighbors, &N3, &N17)
		N19.neighbors = append(N19.neighbors, &N2, &N3)

		N1.name, N2.name, N3.name, N4.name, N5.name, N6.name = "N1", "N2", "N3", "N4", "N5", "N6"
		N7.name, N8.name, N9.name, N10.name, N11.name, N12.name = "N7", "N8", "N9", "N10", "N11", "N12"
		N13.name, N14.name, N15.name, N16.name, N17.name, N18.name, N19.name = "N13", "N14", "N15", "N16", "N17", "N18", "N19"

		maze = append(maze, &N1, &N2, &N3, &N4, &N5, &N6, &N7, &N8, &N9, &N10, &N11, &N12, &N13, &N14, &N15, &N16, &N17, &N18, &N19)
		fmt.Println(len(maze))
	} else {
		//light get off N17,N18

		var N1, N2, N3, N4, N5, N6, N7, N8, N9, N10, N11, N12, N13, N14, N15, N16, N17 Node

		N1.location.x, N1.location.y = 0.5, 13.0
		N2.location.x, N2.location.y = 10.5, 13.0
		N3.location.x, N3.location.y = 13.5, 11.0
		N4.location.x, N4.location.y = 11.5, 0.0
		N5.location.x, N5.location.y = 9.5, 0.0
		N6.location.x, N6.location.y = 0.5, 6.0
		N7.location.x, N7.location.y = 0.5, 1.0
		N8.location.x, N8.location.y = 0.0, 1.0
		N9.location.x, N9.location.y = 0.0, 0.0
		N10.location.x, N10.location.y = 1.5, 6.0
		N11.location.x, N11.location.y = 1.5, 2.0
		N12.location.x, N12.location.y = 9.5, 2.0
		N13.location.x, N13.location.y = 16.5, 0.0
		N14.location.x, N14.location.y = 16.5, 9.0
		N15.location.x, N15.location.y = 11.5, 10.0
		N16.location.x, N16.location.y = 13.5, 9.0

		//N17.location.x, N17.location.y = 9.0, 10.0
		//N18.location.x, N18.location.y = 9.0, 11.0
		N17.location.x, N17.location.y = 13.5, 13.0

		N1.neighbors = append(N1.neighbors, &N6)
		N2.neighbors = append(N2.neighbors, &N17)
		//	N3.neighbors = append(N3.neighbors, &N16, &N18, &N19)
		N3.neighbors = append(N3.neighbors, &N16, &N17)
		N4.neighbors = append(N4.neighbors, &N5, &N13, &N15)
		N5.neighbors = append(N5.neighbors, &N4, &N9, &N12)
		N6.neighbors = append(N6.neighbors, &N1, &N10, &N7)
		N7.neighbors = append(N7.neighbors, &N6, &N8)
		N8.neighbors = append(N8.neighbors, &N7, &N9)
		N9.neighbors = append(N9.neighbors, &N5, &N8)
		N10.neighbors = append(N10.neighbors, &N6, &N11)
		N11.neighbors = append(N11.neighbors, &N10, &N12)
		N12.neighbors = append(N12.neighbors, &N5, &N11)
		N13.neighbors = append(N13.neighbors, &N4, &N14)
		N14.neighbors = append(N14.neighbors, &N13, &N16)
		N15.neighbors = append(N15.neighbors, &N4)
		N16.neighbors = append(N16.neighbors, &N3, &N14)
		//N17.neighbors = append(N17.neighbors, &N15, &N18)
		//N18.neighbors = append(N18.neighbors, &N3, &N17)
		N17.neighbors = append(N17.neighbors, &N2, &N3)

		N1.name, N2.name, N3.name, N4.name, N5.name, N6.name = "N1", "N2", "N3", "N4", "N5", "N6"
		N7.name, N8.name, N9.name, N10.name, N11.name, N12.name = "N7", "N8", "N9", "N10", "N11", "N12"
		N13.name, N14.name, N15.name, N16.name, N17.name = "N13", "N14", "N15", "N16", "N17"
		maze = append(maze, &N1, &N2, &N3, &N4, &N5, &N6, &N7, &N8, &N9, &N10, &N11, &N12, &N13, &N14, &N15, &N16, &N17)

	}

	return maze
	/*
		fmt.Println("address of N1", maze[0])                            // &{[0xc000086180] 0 {0 0} N1}
		fmt.Println("attributes of N1", *maze[0])                        //{[0xc000086180] 0 {0 0} N1}
		fmt.Println("name of N1", (*maze[0]).name)                       //"N1"
		fmt.Println(" address of neighbors of N1", (*maze[0]).neighbors) //[&N6] [0xc0000980f0]
		address := (*maze[0]).neighbors[0]                               //&N6
		fmt.Println("address of N6", address)                            //&N6
		fmt.Println("name of N6", (*address).name)                       //"N6"

		Check if the intialization is right
		For each node
		for i := range maze {
			//address of neighbors of Ni
			neighbors := (*maze[i]).neighbors
			var neighbornames []string
			for j := range neighbors {
				neighbornames = append(neighbornames, (*neighbors[j]).name)
			}
			fmt.Println("location and neighbor names of each node", (*maze[i]).name, (*maze[i]).location, neighbornames)

		}
	*/
}

//MazeEvolve takes in maze as input and return the flow quantity matrix
func MazeEvolve(maze Maze, numgen, nsteps int, t float64) []Matrix {
	N := len(maze)
	// Initialize Pij, Qij, Dij
	P := make([][]float64, numgen+1)
	for i := range P {
		P[i] = make([]float64, N)
	}

	Q := make([]Matrix, numgen+1)
	Q[0] = InitializeFirstMatrix(N, 0.0)

	D := make([]Matrix, numgen+1)

	//Initialize D[0]_ij=0.5, d= D[0]
	D[0] = InitializeFirstMatrix(N, 0.5)

	//Compute the length of node i and its neighbor
	L := ComputeLengthMatrix(maze)

	//Evolve
	for n := 1; n <= numgen; n++ {

		//Compute Pij
		P[n], _ = ComputeP(D[n-1], L)
		//Compute Qij
		Q[n] = ComputeQ(D[n-1], L, P[n])

		//Check if sum_{j}Qij=0 (i!=0/1)
		CheckQ(Q[n])

		D[n] = CalculateConductivity(Q[n], t, nsteps)

		count := 0

		for i := range Q[n] {
			for j := range Q[n][i] {
				if math.Abs(Q[n][i][j]-Q[n-1][i][j]) < 0.00001 {
					count++
				}
			}
		}
		if count == len(Q[0])*len(Q[0][0]) {
			break

		}
	}
	return Q
}

func CheckQ(Q Matrix) {
	for i := range Q {
		s := 0.0
		for j := range Q[0] {
			s += Q[i][j]
		}
		if i != 1 && i != 0 && math.Abs(s) > 0.00001 {
			fmt.Printf("sum flux for node %d is %f not 0\n", i, s)
		}
	}
}
func ComputeQ(d, L Matrix, p []float64) Matrix {
	N := len(d)

	//initialize q
	q := make(Matrix, N)
	for r := range q {
		q[r] = make([]float64, N)
	}

	for i := 0; i <= N-1; i++ {
		for j := 0; j <= N-1; j++ {
			if L[i][j] != 0 {
				q[i][j] = (d[i][j] * (p[i] - p[j])) / L[i][j]
			}
		}
	}

	return q
}

//Check if A is positive and linear dependent?
//CalculateP takes in D and L matrix, P2=0. Solve the remaining n-1 P by n equations. Returning p matrix is a list of float64 variables.
func ComputeP(d, L Matrix) ([]float64, error) {

	//solve linear equations by guassian GaussianElimination : AX=b
	//Initialize the faction matrix dij/lij, augment A by adding the b terms.

	A := InitializePCoefficient(d, L)
	// delete the second column because p2=0
	for i := range A {
		A[i] = append(A[i][:1], A[i][2:]...)
	}

	N := len(A)
	M := len(A[0]) //N-1
	// fmt.Println("N", N)
	// fmt.Println("M", M)
	//For all rows
	i := 0
	k := 0
	for i <= N-1 && k <= M-1 {
		/* Find the k-th pivot:*/
		iMax := Argmax(i, k, A)

		if A[iMax][k] == 0 {
			/*no pivot in this column*/
			k++
		} else {
			A[i], A[iMax] = A[iMax], A[i]

			//for all rows below pivot, change their values
			for r := i + 1; r <= N-1; r++ {
				f := -A[r][k] / A[i][k]
				//for all remaining nonzero elements in current row
				for c := k + 1; c <= M-1; c++ {
					A[r][c] += f * A[i][c]
				}
				//Fill 0 with the lower part of the pivot column
				A[r][k] = 0
			}
		}
		/*Increase pivot row and column*/
		//fmt.Println("i,Ai", i, A[i])
		i++
		k++
	}

	//retrive p from the upper triangular matrix
	p := make([]float64, N)
	for i := N - 2; i >= 0; i-- {
		p[i] = A[i][M-1]
		for j := i + 1; j <= M-1; j++ {
			p[i] -= A[i][j] * p[j]
		}
		p[i] /= A[i][i]
	}

	for i := N - 1; i >= 2; i-- {
		p[i] = p[i-1]
	}
	p[1] = 0.0
	return p, nil
}

//i_max := argmax (i = h ... m, abs(A[i, k]))*/
func Argmax(i, k int, A Matrix) int {
	N := len(A)
	max := A[i][k]
	imax := i
	for r := i + 1; r <= N-1; r++ {
		if temp := A[r][k]; temp > max {
			max = temp
			imax = r
		}
	}
	return imax
}

func InitializePCoefficient(d, L Matrix) Matrix {
	N := len(d)
	M := N + 1 //number of columns
	//The dimention of A is N*(N+1),the extra column is for augment b terms.
	A := make(Matrix, N)
	for r := range A {
		A[r] = make([]float64, M)
	}

	for i := 0; i <= N-1; i++ {
		for j := 0; j <= M-1; j++ {
			//For Dij/Lij
			if j <= N-1 {
				if L[i][j] != 0 && i != j {
					A[i][j] = d[i][j] / L[i][j]
					A[i][i] -= A[i][j]
					//fmt.Println("i,Aii", i, A[i][i])
				}
			} else {
				if i == 0 {
					//For the starting point N1, b=A[0][N] =-1
					A[i][j] = -1
				}
				//For the end point N2, b=b=A[1][N] = =1
				if i == 1 {
					A[i][j] = 1
				}
			}
		}
	}

	return A
}

//Step 10 Calculates the conductivity D of each tube according to Eqs.6 and
//Semi-implicit Euler method (calculation in notability); t = 1s, n = 100
func CalculateConductivity(Q Matrix, t float64, n int) Matrix {
	lenQ := len(Q)
	D := make(Matrix, lenQ)
	for i := 0; i < lenQ; i++ {
		D[i] = make([]float64, lenQ)
	}

	for i := 0; i < lenQ; i++ {
		for j := 0; j < lenQ; j++ {
			D[i][j] = CalculateTubeConductivity(Q[i][j], D[i][j], t, n)
		}
	}

	return D

}

func CalculateTubeConductivity(Qij, Dijn float64, t float64, n int) float64 {
	Dijn1 := 0.0

	for i := 0; i < n; i++ {
		//Dijn1 = ((float64(n)/t)*math.Abs(Qij) + Dijn) / (1 + float64(n)/t)
		Dijn1 = (Dijn1/t + math.Abs(Qij)) / (1.0 + 1.0/t)
	}

	return Dijn1
}

//ComputeLength takes in the maze and returns the length between the distance from the node to its neighbors
func ComputeLengthMatrix(maze Maze) Matrix {
	N := len(maze)

	L := make(Matrix, N)
	for r := range L {
		L[r] = make([]float64, N)
	}

	//Calculate distances of one node to its all neighbors
	for i := range maze {
		//address of neighbors of Ni
		neighbors := (*maze[i]).neighbors
		for j := range neighbors {
			//j is just the order of the neighbor in neighbors. To know its real index in the maze and L
			// we need to extract it from the name, the real index is the indexj-1 N1->0
			indexj, _ := strconv.Atoi((*(neighbors[j])).name[1:])
			indexj--
			//fmt.Println((*(neighbors[0])).name[1:])
			l := ComputeLength(maze, i, j)
			L[i][indexj], L[indexj][i] = l, l
			// fmt.Println("length of node", i+1, "and node", indexj+1, "is", l)
		}
	}
	return L
}

func ComputeLength(maze Maze, i, j int) float64 {
	node := *maze[i]
	nodeneighborj := *(node.neighbors[j])
	return math.Abs(node.location.x-nodeneighborj.location.x) + math.Abs(node.location.y-nodeneighborj.location.y)
}

//Initialize D takes the length of D[0] and returns N*N matrix where value of each entry is 0.5
func InitializeFirstMatrix(N int, value float64) Matrix {
	d := make(Matrix, N)
	for r := range d {
		d[r] = make([]float64, N)
	}

	for r := range d {
		for c := range d[r] {
			d[r][c] = value
		}
	}
	return d
}
