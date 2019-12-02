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
	// dampT := 0.3
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
	// emptyboard := InitializeBoard(row, col)
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

	numGens := 20000
	numInterval := 20
	fileName := "corner"
	fmt.Println("All command line arguments read successfully.")
	//Return boards to plot.

	start1 := time.Now()
	quickboards := SimulateSlimeMold(matrix0, numGens, numInterval, sensorArmLength, sensorAngle, sensorDiagonalL, depT, dampT, filterT, WL, WT, WN, CN, CL, dampN, filterN, RT, ET)
	elapsed1 := time.Since(start1)
	log.Printf("model takes %s\n", elapsed1)

	start2 := time.Now()
	imagefile := DrawGameBoards(quickboards, 1, CN)
	ImagesToGIF(imagefile, fileName)
	elapsed2 := time.Since(start2)
	log.Printf("drawing takes %s\n", elapsed2)

}

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
			/*
				for j := range currBoardCopy[0] {
					if math.Abs(currBoardCopy[0][j].trailChemo) > 0.0001 || math.Abs(currBoardCopy[0][j].foodChemo) > 0.0001 {
						fmt.Printf("boundary 0,%d is %f, %f, not 0\n", j, currBoardCopy[0][j].trailChemo, currBoardCopy[0][j].foodChemo)
					}
					if math.Abs(currBoardCopy[numRows-1][j].trailChemo) > 0.0001 || math.Abs(currBoardCopy[numRows-1][j].foodChemo) > 0.0001 {
						fmt.Printf("boundary numRows-1,%d is %f, %f, not 0\n", j, currBoardCopy[numRows-1][j].trailChemo, currBoardCopy[numRows-1][j].foodChemo)
					}
				}

				for i := range currBoardCopy {
					if math.Abs(currBoardCopy[i][0].trailChemo) > 0.0001 || math.Abs(currBoardCopy[i][0].foodChemo) > 0.0001 {
						fmt.Printf("boundary %d,0 is %f, %f,not 0\n", i, currBoardCopy[i][0].trailChemo, currBoardCopy[i][0].foodChemo)
					}
					if math.Abs(currBoardCopy[i][numCols-1].trailChemo) > 0.0001 || math.Abs(currBoardCopy[i][numCols-1].foodChemo) > 0.0001 {
						fmt.Printf("boundary %d,numCols-1 is %f ,%f, not 0\n", i, currBoardCopy[i][numCols-1].trailChemo, currBoardCopy[i][numCols-1].foodChemo)
					}
				}
			*/
		}
	}

	return quickboards
}
