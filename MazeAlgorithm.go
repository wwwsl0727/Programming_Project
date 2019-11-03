package main

import (
	"fmt"
)

type matrix [][]float64

//based on maze, initialize matrix L and number of nodes N

type orderpair struct{
  x,y float64
}

type node struct {
  neighbor []*node
  pressure float64
  location orderpair
  name     string
}



func main() {
	fmt.Println("Maze Algorithm")
	N := 0                     //total number of nodes
	L := make(Matrix, N, N)    // a matrix of length, in which L[i][j] is the length of the tube connecting node Ni and Nj
	Q := make([]Matrix, N, N)    // a slice of matrix to record every round of current moment flux in each tube, in which Q[][i][j] is the flux through the tube from node Ni to Nj;
	D := make([]Matrix, N, N)    // a slice of matrix to record every round of conductivities, in which D[][i][j] is the conductivity of tube L[][i][j];
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



//Step 4 Set iterate number count = 0, the flux in each tube Q[i][j] = 0,(∀i,j = 1,2,...,N)
//and Qpre[i][j] = Q[i][j] , the flux change in each tube Q cha[i][j] = 0, (∀i, j = 1, 2, . . . , N),
//the conductivities of each tube in the maze D[j][i] = D[i][j] = C[i][j] (where Cij ranges from 0 to 1),
//and the pressure of each node Pi = 0;


//Step 5 Set count = count + 1 and calculate the pressure P at each node according to Eqs. 5;
//Take i, j, D[i][j], L[i][j], pi, pj, Q[i][j]
func calculatePressue() float64 {

}



//step 6 Calculate the flux Q in each tube according to Equation 1.
func CalculateFlux(D, L , pi, pj float64) float64 {
  return D/L * (pi - pj)
}
