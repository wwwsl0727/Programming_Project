package main

import (
	"fmt"
)

type Maze []*Node

type Node struct {
	neighbors []*Node
	pressure  float64
	location  OrderedPair
	name      string
}

type OrderedPair struct {
	x, y float64
}

func main() {
	fmt.Println("Maze!")

	//Initialize mazes.
	var N1, N2, N3, N4, N5, N6, N7, N8, N9, N10, N11, N12, N13, N14, N15, N16, N17, N18, N19 Node

	N1.location.x, N1.location.y = 0.0, 0.0
	N2.location.x, N2.location.y = 10.0, 0.0
	N3.location.x, N3.location.y = 13.0, 2.0
	N4.location.x, N4.location.y = 11.0, 13.0
	N5.location.x, N5.location.y = 9.0, 13.0
	N6.location.x, N6.location.y = 0.0, 7.0
	N7.location.x, N7.location.y = 0.0, 12.0
	N8.location.x, N8.location.y = -0.5, 12.0
	N9.location.x, N9.location.y = -0.5, 13.0
	N10.location.x, N10.location.y = 1.0, 7.0
	N11.location.x, N11.location.y = 1.0, 11.0
	N12.location.x, N12.location.y = 9.0, 11.0
	N13.location.x, N13.location.y = 16.0, 13.0
	N14.location.x, N14.location.y = 16.0, 4.0
	N15.location.x, N15.location.y = 11.0, 3.0
	N16.location.x, N16.location.y = 13.0, 4.0
	N17.location.x, N17.location.y = 8.5, 3.0
	N18.location.x, N18.location.y = 8.5, 2.0
	N19.location.x, N19.location.y = 13.0, 0.0

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
	var maze Maze
	maze = append(maze, &N1, &N2, &N3, &N4, &N5, &N6, &N7, &N8, &N9, &N10, &N11, &N12, &N13, &N14, &N15, &N16, &N17, &N18, &N19)
	fmt.Println("address of N1", maze[0])                            // &{[0xc000086180] 0 {0 0} N1}
	fmt.Println("attributes of N1", *maze[0])                        //{[0xc000086180] 0 {0 0} N1}
	fmt.Println("name of N1", (*maze[0]).name)                       //"N1"
	fmt.Println(" address of neighbors of N1", (*maze[0]).neighbors) //[&N6] [0xc0000980f0]
	address := (*maze[0]).neighbors[0]                               //&N6
	fmt.Println("address of N6", address)                            //&N6
	fmt.Println("name of N6", (*address).name)                       //"N6"

	//Check if the intialization is right
	//For each node
	// for i := range maze {
	// 	//address of neighbors of Ni
	// 	neighbors := (*maze[i]).neighbors
	// 	var neighbornames []string
	// 	for j := range neighbors {
	// 		neighbornames = append(neighbornames, (*neighbors[j]).name)
	// 	}
	// 	fmt.Println("location and neighbor names of each node", (*maze[i]).name, (*maze[i]).location, neighbornames)
	//
	// }

	//Initialize Pij, Qij, Dij
	//Compute Pij
	//Compute Qij
	//Compute Dij
}
