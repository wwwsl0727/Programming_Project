package main

import (
	"fmt"
)

type matrix [][]float64

//based on maze, initialize matrix L and number of nodes N

func main() {
	fmt.Println("Maze Algorithm")
	N := 0                     //total number of nodes
	L := make(Matrix, N, N)    // a matrix of length, in which L[i][j] is the length of the tube connecting node Ni and Nj
	Q := make(Matrix, N, N)    // a matrix to record current moment flux in each tube, in which Q[i][j] is the flux through the tube from node Ni to Nj;
	Qpre := make(Matrix, N, N) // a matrix to record previous moment flux in each tube;
	Qcha := make(Matrix, N, N) // a matrix to record change of flux in each tube;
	D := make(Matrix, N, N)    // a matrix to record conductivities, in which D[i][j] is the conductivity of tube L[i][j];
	P := make([]float64, N)    // pressure at each node; Pi = 0 if Ni is the exit node
}

func removeNodeAndTube (L matrix, N int) matrix, int{
  for i := range L{
    if NodeHasXTube(L, i, x) == true {
      N = N - 1
      L = RemoveOneNode(L, i)
      }
    }
    return L, N
}

// NodeHasXTube takes a length matrix L, an node index i, and number of tubes x.
// Returns true if matrix[i] has n numbers. (node i has n tubes)
func NodeHasXTube(L matrix, i int, x int) bool {
  count := 0
  for k := range L[i]{
    if L[i][k] != nil {
      count ++
    }
  }
  if count == x{
    return true
  }
}

//Remove one node from the matrix. Delete matrix[i] and matrix[][i]
func RemoveOneNode(L matrix, i int) matrix {

}
