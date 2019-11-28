package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
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

type Matrix [][]float64

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
	// fmt.Println("address of N1", maze[0])                            // &{[0xc000086180] 0 {0 0} N1}
	// fmt.Println("attributes of N1", *maze[0])                        //{[0xc000086180] 0 {0 0} N1}
	// fmt.Println("name of N1", (*maze[0]).name)                       //"N1"
	// fmt.Println(" address of neighbors of N1", (*maze[0]).neighbors) //[&N6] [0xc0000980f0]
	// address := (*maze[0]).neighbors[0]                               //&N6
	// fmt.Println("address of N6", address)                            //&N6
	// fmt.Println("name of N6", (*address).name)                       //"N6"

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

	numgen := 10

	nsteps := 10000 // This is for calculating conductivity

	t := 1.0

	filename := "visualization"
	rand.Seed(time.Now().UnixNano())
	Q := MazeEvolve(maze, numgen, nsteps, t)

	// for i := 0; i < len(Q); i++ {
	// 	for j := 0; j < len(Q[i]); j++ {
	// 		for k := 0; k < len(Q[i][j]); k++ {
	// 			fmt.Print(Q[i][j][k], " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }
	imageList := DrawMazes(maze, Q, numgen, 14.0)
	ImagesToGIF(imageList, filename)

	// fmt.Println(MazeEvolve(maze, numgen, nsteps, t))

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

		D[n] = CalculateConductivity(Q[n], t, nsteps)
		count := 0

		for i := range Q[n] {
			for j := range Q[n][i] {
				if math.Abs(Q[n][i][j]-Q[n-1][i][j]) < 10^(-5) {
					count++
				}
			}
		}
		if count == len(Q)*len(Q[0]) {
			break

		}
	}
	return Q
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
		Dijn1 = ((float64(n)/t)*math.Abs(Qij) + Dijn) / (1 + float64(n)/t)
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
