package main

/*
func TestComputeP(t *testing.T) {
	N := 3
	d := [][]float64{{0.0, 3.0, -1.0},
		{3.0, 0.0, 1.0},
		{-1.0, 1.0, 0.0}}
	L := InitializeFirstMatrix(N, 1.0)
	P, _ := ComputeP(d, L)
	expectedP := []float64{0.0, 0.0, 1.0}
	for i := range P {
		if P[i] != expectedP[i] {
			fmt.Println(expectedP[0])
			fmt.Printf("%d as %f doesn't match %f", i, P[i], expectedP[i])
		}
	}
}
*/
/*
func TestInitializePCoefficient(t *testing.T) {
	N := 3
	d := InitializeFirstMatrix(N, 1.0)
	L := [][]float64{{0.0, 3.0, 0.0},
		{3.0, 0.0, 4.0},
		{0.0, 4.0, 0.0}}
	A := InitializePCoefficient(d, L)
	expectedA := [][]float64{
		{-1.0 / 3.0, 1.0 / 3.0, 0.0, -1.0},
		{1.0 / 3.0, -(1.0/3.0 + 1.0/4.0), 1.0 / 4.0, 1.0},
		{0.0, 1.0 / 4.0, -1.0 / 4.0, 0.0},
	}
	for i := range A {
		for j := range A[0] {
			if A[i][j] != expectedA[i][j] {
				fmt.Printf("%d,%d as %f doesn't match %f", i, j, A[i][j], expectedA[i][j])
			}
		}
	}
}
*/
