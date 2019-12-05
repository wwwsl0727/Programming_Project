package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 2*2 plane, each box represents one grid
type multiAgentMatrix [][]box

type box struct {
	IsFood, IsAgent, haslight bool
	foodChemo                 float64
	trailChemo                float64
	light                     float64
	agent                     Agent
}

type Agent struct {
	motionCounter   int     // initial 0
	direction       float64 // direction could be the multiple of pi/4
	sensorLength    int     // 7 in 50% agent case
	sensorDiagonalL float64 // float64(5)*math.Sqrt(2) for sensor in diagnol of axix
	isMoved         bool    //default false, if the agent in
}

func main() {

	//command summary:
	// ./multiAgent food
	// ./multiAgent light int int
	// ./multiAgent wind int
	// ./multiAgent normal string
	rand.Seed(time.Now().UnixNano())

	//Below is initialization for 50% case, three food spots.
	//Different intialization will change variables below accordingly.
	row := 200
	col := 200
	sensorArmLength := 7
	sensorAngle := math.Pi / float64(4)
	sensorDiagonalL := float64(5) * math.Sqrt(2)
	depT := 5.0  //The quantity of trail deposited by an agent
	dampT := 0.1 //Diffusion damping factor of trail
	// filter for trail 3x3
	filterT := 3

	WL := 1.0    //WL: Weight of light value sensed by an agent’s sensor
	WT := 0.4    //WT: Weight of trail value
	WN := 1 - WT //WN: Weight of nutrient value
	CN := 10.0   //The Chemo-Nutrient concentration of food
	CL := 10.0   // Light concentration
	dampN := 0.2 //Diffusion damping factor of chemo- nutrient
	// filter for nutrient 5x5
	filterN := 5
	//motion counter >= RT, reproduction; motion counter < ET, elimation
	RT := 15
	ET := -10
	matrix0 := InitializeBoard(row, col) // Used to pass to later simulation after initialization
	fmt.Println("please input condition(normal), situation(half/corner), numGens,numIntervals,filename")
	//Different Initilization based on command line
	condition := os.Args[1]
	if condition == "food" { //50% mold, two good foods, two bad foods with chemo 0.
		matrix0 = intializeFoodBoard(matrix0, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)

	} else if condition == "light" {
		x, err2 := strconv.Atoi(os.Args[2]) // light center x index
		if err2 != nil {
			panic("Issue in read light x index")
		}
		y, err3 := strconv.Atoi(os.Args[3]) // light center y index
		if err3 != nil {
			panic("Issue in read light y index")
		}
		matrix0 = intializeLightBoard(matrix0, row, col, sensorArmLength, x, y, sensorDiagonalL, sensorAngle, CN, CL)

	} else if condition == "wind" {
		windlevel, err2 := strconv.Atoi(os.Args[2]) // wind level is 0 ~ 10
		if err2 != nil {
			panic("Issue in read wind level")
		}
		RT += windlevel //originally 15, make it harder to reproduce
		ET += windlevel //originally -10, make it easy to die
		matrix0 = intializeHalfBoard(matrix0, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)

	} else if condition == "normal" {
		situation := os.Args[2]
		if situation == "half" {
			//half of the board has mold and three food spots as triangle
			matrix0 = intializeHalfBoard(matrix0, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)

			// all molds start in a corner of board
		} else if situation == "corner" {
			//Need to adjust certain variable such as WT
			CN = 20
			filterN = 13
			WN = 0.8
			WT = 0.2
			matrix0 = intializeCornerBoard(matrix0, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)
			//******* need to change other factors
		}
	} else {
		panic("wrong condition input!")
	}

	numGens_input := os.Args[3]
	numGens, _ := strconv.Atoi(numGens_input) // number of generations
	numInterval_input := os.Args[4]
	numInterval, _ := strconv.Atoi(numInterval_input) // the interval to draw a picture
	fileName := os.Args[5]              // the filename of the GIF.
	fmt.Println("All command line arguments read successfully.")
	//Return boards to plot.

	//Start simulation
	start1 := time.Now()
	quickboards := SimulateSlimeMold(matrix0, numGens, numInterval, sensorArmLength, sensorAngle, sensorDiagonalL, depT, dampT, filterT, WL, WT, WN, CN, CL, dampN, filterN, RT, ET)
	elapsed1 := time.Since(start1)
	log.Printf("model takes %s\n", elapsed1)

	//Draw boards.
	start2 := time.Now()
	imagefile := DrawGameBoards(quickboards, 1, CN)
	ImagesToGIF(imagefile, fileName)
	elapsed2 := time.Since(start2)
	log.Printf("drawing takes %s\n", elapsed2)

}


//Simulation takes all parameters and store a matrix every numInterval generations.
func SimulateSlimeMold(matrix0 multiAgentMatrix,
	numGens int,
	numInterval int,
	sensorArmLength int,
	sensorAngle float64,
	sensorDiagonalL float64,
	depT float64,
	dampT float64,
	filterT int,
	WL float64,
	WT float64,
	WN float64,
	CN float64,
	CL float64,
	dampN float64,
	filterN int,
	RT int,
	ET int) []multiAgentMatrix {
	numRows := len(matrix0)
	numCols := len(matrix0[0])
	numBoards := numGens/numInterval + 1

	//The boards to return for GIF.
	quickboards := make([]multiAgentMatrix, numBoards)
	//Initialize every element of quickboards
	for i := range quickboards {
		quickboards[i] = InitializeBoard(numRows, numCols)
	}
	quickboards[0] = matrix0

	currBoard := CopyBoard(matrix0)

	//Initialize index array,this is used to update agent in a shuffle order later
	//1,2,...len(board)-2
	rowIndices := make([]int, numRows-2)
	for i := 1; i <= numRows-2; i++ {
		rowIndices[i-1] = i
	}
	colIndices := make([]int, numCols-2)
	for i := 1; i <= numCols-2; i++ {
		colIndices[i-1] = i
	}

	//Append the matrix to the returning slice of matrixes every numInterval iteration.
	for n := 1; n < numGens+1; n++ {

		boardCopy := CopyBoard(currBoard)
		//Update board based on the copy
		currBoard.UpdateBoard(boardCopy, sensorArmLength, sensorAngle,
			sensorDiagonalL, depT, dampT,
			filterN, WL, WT, WN,
			CN, CL, dampN, filterT, RT, ET, rowIndices, colIndices)

		if n%(numInterval) == 0 {
			currBoardCopy := CopyBoard(currBoard)
			quickboards[n/numInterval] = currBoardCopy

		}
	}

	return quickboards
}
