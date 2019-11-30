package main

import (
	"fmt"
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

	numGens := 10000

	emptyboard := InitializeBoard(row, col)
	matrix0 := InitializeBoard(row, col) // Used to pass to later simulation after initialization

	//Different Initilization based on command line
	condition := os.Args[1]
	if condition == "food" { //50% mold, two good foods, two bad foods with chemo 0.
		matrix0 = intializeFoodBoard(emptyboard, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)

	} else if condition == "light" {
		x, err2 := strconv.Atoi(os.Args[2]) // light center x index
		if err2 != nil {
			panic("Issue in read light x index")
		}
		y, err3 := strconv.Atoi(os.Args[3]) // light center y index
		if err3 != nil {
			panic("Issue in read light y index")
		}
		matrix0 = intializeLightBoard(emptyboard, row, col, sensorArmLength, x, y, sensorDiagonalL, sensorAngle, CN, CL)

	} else if condition == "wind" {
		windlevel, err2 := strconv.Atoi(os.Args[2]) // wind level is 0 ~ 10
		if err2 != nil {
			panic("Issue in read wind level")
		}
		RT += windlevel //originally 15, make it harder to reproduce
		ET += windlevel //originally -10, make it easy to die

	} else if condition == "normal" {
		situation := os.Args[2]
		if situation == "half" {
			//half of the board has mold and three food spots as triangle
			matrix0 = intializeHalfBoard(emptyboard, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)

			// all molds start in a corner of board
		} else if situation == "corner" {
			//Need to adjust certain variable such as WT
			matrix0 = intializeCornerBoard(emptyboard, row, col, sensorArmLength, sensorDiagonalL, sensorAngle, CN)
			//******* need to change other factors
		}
	} else {
		panic("wrong condition input!")
	}
	fmt.Println("All command line arguments read successfully.")

	//Return boards to plot.
	boards := make([]multiAgentMatrix, numGens+1)
	for i := range boards {
		boards[i] = InitializeBoard(row, col)
	}
	boards[0] = matrix0

	//start simulation
	for n := 0; n <= numGens-1; n++ {

		currBoard := CopyBoard(boards[n])
		//For boundary,both the foodChemo and the trailChemo is 0.
		currBoard = UpdateChemo(filterN, dampN, boards[n], "food")
		//fmt.Println(currBoard)
		currBoard = UpdateChemo(filterT, dampT, boards[n], "trial")

		//for each agent, update each agent
		for r := range currBoard {
			for c := range currBoard[0] {
				//sense direction and change direction to the higher chemo direction.
				if currBoard[r][c].IsAgent {
					agentDirection := currBoard.SynthesisComparator(r, c, WT, WN, WL, sensorAngle)

					//direction 0,+- pi/2; +- pi; +- 3pi/2; +-2pi... / pi/4 = 0;2;4;6... a=L;else a=DiagonalL
					flag := int(4 * agentDirection / math.Pi)
					var forwardx int
					var forwardy int

					var a float64
					if flag%2 == 0 {
						a = float64(sensorArmLength)
					} else {
						a = sensorDiagonalL
					}

					forwardx = r + int(a*math.Cos(agentDirection))
					forwardy = c + int(a*math.Sin(agentDirection))

					currCell := &currBoard[r][c]
					forwardcell := &currBoard[forwardx][forwardy]
					//If the forward direction is occupied, change direction randomly and motionCounter--
					if forwardcell.IsAgent == true {
						currCell.agent.direction = float64(rand.Intn(8)+1) * sensorAngle
						currCell.agent.motionCounter--
						if currCell.agent.motionCounter < ET {
							//currentCell die
							currCell.IsAgent = false
							// currCell.agent = nil
						}
					} else {
						//If the forward direction is not occupied, move to that direction and leave trail in the new cell
						forwardcell.IsAgent = true
						forwardcell.trailChemo += depT
						var forwardAgent Agent
						forwardAgent.motionCounter = currCell.agent.motionCounter + 1
						//rotate to the direction with the higher sense value
						forwardAgent.direction = currBoard.SynthesisComparator(forwardx, forwardy, WT, WN, WL, sensorAngle)
						forwardAgent.sensorDiagonalL = sensorDiagonalL
						forwardAgent.sensorLength = sensorArmLength
						forwardcell.agent = forwardAgent

						//Check motionCounter>RT
						if forwardcell.agent.motionCounter > RT {
							//a new cell born in the father cell
							currCell.agent.direction = float64(rand.Intn(8)+1) * sensorAngle
							currCell.agent.motionCounter = 0
							currCell.agent.sensorDiagonalL = sensorDiagonalL
							currCell.agent.sensorLength = sensorArmLength
						} else {
							//there is no new cell in the father cell.
							(*currCell).IsAgent = false
							// (*currCell).agent = nil
						}

					}
				}

			}

		}
		boards[n+1] = currBoard

	}
	// for i := 9; i <= 11; i++ {
	// 	for j := 99; j <= 101; j++ {
	// 		fmt.Println(boards[0][i][j].IsFood)
	// 	}
	// }
	//fmt.Println(boards[1])

	// quickboards := make([]multiAgentMatrix, 0)
	// for i := 0; i < len(boards); i += 100 {
	// 	quickboards = append(quickboards, boards[i])
	// 	fmt.Println(len(quickboards))
	// }
	// fmt.Println(len(quickboards))
	//
	// imagefile := DrawGameBoards(quickboards, 1, CN)
	// ImagesToGIF(imagefile, "Multiagent_GIF")
}

//50% mold, two good foods, two bad foods with chemo 0.
func intializeFoodBoard(matrix0 multiAgentMatrix, row, col, sensorArmLength int, sensorDiagonalL, sensorAngle, CN float64) multiAgentMatrix {
	for i := range matrix0 {
		//for all boxes, intial foodChemo & trailChemo are 0.
		//IsFood, IsAgent are false
		row := make([]box, col)
		row = GenerateAgent(row, sensorArmLength, sensorDiagonalL, sensorAngle)
		matrix0[i] = row
	}
	//Bad food center is 50,50.
	for i := 49; i <= 51; i++ {
		for j := 49; j <= 51; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = 0.0 //Bad food
		}
	}

	//Bad food center is 150,50.
	for i := 149; i <= 151; i++ {
		for j := 49; j <= 51; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = 0.0 //Bad food
		}
	}

	//Good food center is 50,150.
	for i := 49; i <= 51; i++ {
		for j := 149; j <= 151; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //Good food
		}
	}

	//Bad food center is 150,150.
	for i := 149; i <= 151; i++ {
		for j := 149; j <= 151; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //good food
		}
	}
	return matrix0
}

//50% mold, three foods, light square depends on x and y
func intializeLightBoard(matrix0 multiAgentMatrix, row, col, sensorArmLength, x, y int, sensorDiagonalL, sensorAngle, CN, CL float64) multiAgentMatrix {
	for i := range matrix0 {
		//for all boxes, intial foodChemo & trailChemo are 0.
		//IsFood, IsAgent are false
		row := make([]box, col)
		row = GenerateAgent(row, sensorArmLength, sensorDiagonalL, sensorAngle)
		matrix0[i] = row
	}

	//The center is 10,100
	for i := 9; i <= 11; i++ {
		for j := 99; j <= 101; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,10
	for i := 189; i <= 191; i++ {
		for j := 9; j <= 11; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,190
	for i := 189; i <= 191; i++ {
		for j := 189; j <= 191; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//Given the center of light, the boxes in light source has stronger light concentration
	addlight(matrix0, x, y, CL)

	return matrix0
}

//Add light concentration to a square of length 9
func addlight(matrix0 multiAgentMatrix, x, y int, CL float64) {
	for i := x - 4; i <= x+4; i++ {
		for j := y - 4; i <= y+4; i++ {
			if InField(200, 200, i, j) {
				matrix0[i][j].haslight = true
				matrix0[i][j].light = CL
			}
		}
	}
}

func intializeHalfBoard(matrix0 multiAgentMatrix, row, col, sensorArmLength int, sensorDiagonalL, sensorAngle, CN float64) multiAgentMatrix {
	//make the matrix which has 50% of its box has agent

	for i := range matrix0 {
		//for all boxes, intial foodChemo & trailChemo are 0.
		//IsFood, IsAgent are false
		row := make([]box, col)
		row = GenerateAgent(row, sensorArmLength, sensorDiagonalL, sensorAngle)
		matrix0[i] = row
	}

	//The center is 10,100
	// for i := 9; i <= 11; i++ {
	// 	for j := 99; j <= 101; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //The center is 190,10
	// for i := 189; i <= 191; i++ {
	// 	for j := 9; j <= 11; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //The center is 190,190
	// for i := 189; i <= 191; i++ {
	// 	for j := 189; j <= 191; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }

	//The center is 10,100
	for i := 99; i <= 101; i++ {
		for j := 9; j <= 11; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,10
	for i := 9; i <= 11; i++ {
		for j := 189; j <= 191; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,190
	for i := 189; i <= 191; i++ {
		for j := 189; j <= 191; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}
	return matrix0
}

////////////////////////// Shili's part
func intializeCornerBoard(matrix0 multiAgentMatrix, row, col, sensorArmLength int, sensorDiagonalL, sensorAngle, CN float64) multiAgentMatrix {
	//put 9 agents in the southern part
	for i := 149; i <= 151; i++ {
		for j := 149; j <= 151; j++ {
			matrix0[i][j].IsAgent = true
			var agent Agent
			agent.direction = float64(rand.Intn(8)+1) * sensorAngle
			agent.motionCounter = 0
			agent.sensorDiagonalL = sensorDiagonalL
			agent.sensorLength = sensorArmLength
			matrix0[i][j].agent = agent
		}
	}
	return matrix0
}

func UpdateChemo(filter int, damp float64, board multiAgentMatrix, category string) multiAgentMatrix {
	board1 := CopyBoard(board)
	//use 5*5 average filter
	board1.AverageFilter(filter, category)
	//damp
	board1.Damp(damp, category)
	return board1
}

func (board multiAgentMatrix) AverageFilter(filter int, category string) {
	numRows := len(board)
	numCols := len(board[0])
	boardCopy := CopyBoard(board)
	for i := range board {
		for j := range board[0] {
			//check the neighbors and get the Average.
			// For boundary,both the foodChemo and the trailChemo is 0.
			//Apply food filters only on box which is not food resource
			if i != 0 && i != numRows-1 && j != 0 && j != numCols-1 {
				if category == "food" {
					if !board[i][j].IsFood {
						board.AverageNeighbor(boardCopy, i, j, filter, category)
					}
				} else if category == "trial" {
					board.AverageNeighbor(boardCopy, i, j, filter, category)
				} else {
					panic("the given category is wrong")
				}
			}
		}
	}
}

func (board multiAgentMatrix) AverageNeighbor(boardCopy multiAgentMatrix, r, c, filter int, category string) {
	numRows := len(board)
	numCols := len(board[0])

	var numNeighbors int
	var sum float64
	interval := filter / 2
	for i := r - interval; i <= r+interval; i++ {
		for j := c - interval; j <= c+interval; j++ {
			if InField(numRows, numCols, i, j) {
				numNeighbors++
				if category == "food" {
					sum += boardCopy[i][j].foodChemo
				} else {
					sum += boardCopy[i][j].trailChemo
				}

			}
		}
	}
	ave := sum / float64(numNeighbors)
	if category == "food" {
		board[r][c].foodChemo = ave
	} else {
		board[r][c].trailChemo = ave
		/*
			if r == 4 && c == 4 {
				fmt.Println(sum)
				fmt.Println(numNeighbors)
			}
		*/
	}
}

func (board multiAgentMatrix) Damp(damp float64, category string) {
	factor := 1.0 - damp
	for i := range board {
		for j := range board[i] {
			//trial value of food source can be changed
			if category == "food" {
				if !board[i][j].IsFood {

					board[i][j].foodChemo *= factor
					/*
						if i == 1 && j == 1 {
							fmt.Println(board[i][j].foodChemo)
						}
					*/
				}
			} else {
				board[i][j].trailChemo *= factor
			}

			/*
				if !board[i][j].IsFood {
					if category == "food" {
						board[i][j].foodChemo = (1 - damp) * board[i][j].foodChemo
					} else {
						board[i][j].trailChemo = (1 - damp) * board[i][j].trailChemo
					}
				}
			*/
		}
	}
}

//GenerateAgent Generate agent for a row. 50% of boxes have agent
func GenerateAgent(row []box, sensorArmLength int, sensorDiagonalL, sensorAngle float64) []box {

	for i := range row {
		n := rand.Intn(2) //50% chance of having an agent in the box
		// randomDirection := rand.Intn(8) + 1 //This create a random direction from 1 - 8
		randomDirection := float64(rand.Intn(8)+1) * sensorAngle //This create a random direction from pi/4 -2pi
		if n == 0 {
			row[i].IsAgent = true
			//create a new agent and points to it
			var oneAgent Agent
			oneAgent.motionCounter = 0
			oneAgent.direction = randomDirection
			oneAgent.sensorLength = sensorArmLength
			oneAgent.sensorDiagonalL = sensorDiagonalL
			row[i].agent = oneAgent
		}
		// if n == 1, no agent
	}
	return row
}

//Compare two sensors measure, change a agent's direction
func (matrix multiAgentMatrix) SynthesisComparator(row, col int, WT, WN, WL, sensorAngle float64) float64 {
	//exceptions
	numRows := len(matrix)
	numCols := len(matrix[0])
	if row < 0 || col < 0 || row >= numRows || col >= numCols {
		panic("the given row and col is out of bound")
	}

	if matrix[row][col].IsAgent == false {
		panic("This box has no Agent!")
	}

	agentDirection := matrix[row][col].agent.direction

	//For sensors on the x and y axis, sensorLength is 7
	L := matrix[row][col].agent.sensorLength

	//For sensors on the diagonal axis of xy plane, in order to get the location as integers, sensorlength is 5*sqart(2)
	DiagonalL := matrix[row][col].agent.sensorDiagonalL

	// The direction of the left,right sensor is based on the current direction of the agent.
	leftx, lefty, rightx, righty := CalculateSensorLocation(agentDirection, sensorAngle, DiagonalL, row, col, L)
	// fmt.Println("leftx", leftx)
	// fmt.Println("lefty", lefty)
	// fmt.Println("rightx", rightx)
	// fmt.Println("righty", righty)
	//Get sample chemoattractant values from sensors.
	// If left sensor and right sensor are all in the matrix
	if InField(numRows, numCols, leftx, lefty) && InField(numRows, numCols, rightx, righty) {
		sensorLeft := matrix[leftx][lefty]
		FL := calculateScore(sensorLeft, WT, WN, WL)
		sensorRight := matrix[rightx][righty]
		FR := calculateScore(sensorRight, WT, WN, WL)
		if FL < FR {
			//rotate right
			matrix[row][col].agent.direction += sensorAngle
		} else if FL > FR {
			matrix[row][col].agent.direction -= sensorAngle
		} else {
			//random choose left and right
			if rand.Intn(2) == 0 {
				matrix[row][col].agent.direction += sensorAngle
			} else {
				matrix[row][col].agent.direction -= sensorAngle
			}
		}
	} else if InField(numRows, numCols, leftx, lefty) == false && InField(numRows, numCols, rightx, righty) {
		// If the left sensor is out of bound while the right sensor is in, turn right
		// fmt.Println("direction before change", matrix[row][col].agent.direction)
		matrix[row][col].agent.direction += sensorAngle
		// fmt.Println("direction changes into", matrix[row][col].agent.direction)
		// fmt.Println("left out of bound")
	} else if InField(numRows, numCols, leftx, lefty) && InField(numRows, numCols, rightx, righty) == false {
		//Otherwise, it turns left
		matrix[row][col].agent.direction -= sensorAngle
	} else {
		//If both side is out of bound, the agent turns back
		matrix[row][col].agent.direction += math.Pi
	}
	return matrix[row][col].agent.direction

}

//CalculateSensorLocation Calculate the location of left and right sensors.
func CalculateSensorLocation(agentDirection, sensorAngle, DiagonalL float64, row, col, L int) (int, int, int, int) {
	//Initialize the location of left, forward and right sensors
	var leftx int
	var lefty int
	var rightx int
	var righty int

	flag := int(agentDirection / sensorAngle)

	leftDirection := agentDirection - sensorAngle
	rightDirection := agentDirection + sensorAngle

	//a is the coefficient for left and right arms; b is the coefficient for forward arms
	//a,b =L/DiagonalL
	var a float64
	//Check whether the forward direction of agent is +-pi/4;+-3pi/4;+-5pi/4;+-7pi/4;;
	//If so, a=L;leftx = row+a*cos(leftDirection); lefty= row+a*sin(leftDirection)
	//rightx = row+a*cos(rightDirection); righty= row+a*sin(rightDirection)
	//Else, exchange a=DiagonalL
	if flag%2 != 0 {
		a = float64(L)
	} else {
		a = DiagonalL
	}
	leftx = row + int(a*math.Cos(leftDirection))
	lefty = col + int(a*math.Sin(leftDirection))
	rightx = row + int(a*math.Cos(rightDirection))
	righty = col + int(a*math.Sin(rightDirection))
	return leftx, lefty, rightx, righty
}

//Calculate a box's score using SV = WT×TV +WN×NV
func calculateScore(B box, WT, WN, WL float64) float64 {
	return WT*B.trailChemo + WN*B.foodChemo - WL*B.light
}

//InField takes a the numRows and numCols of a matrix and i/j indices.  It returns true if (i,j) is a valid entry
//of the board.
func InField(numRows, numCols, i, j int) bool {
	if i < 0 || j < 0 || i >= numRows || j >= numCols {
		return false
	}
	// if we make it here, we are in the field.
	return true
}

func InitializeBoard(row, col int) multiAgentMatrix {
	board := make(multiAgentMatrix, row)
	for i := range board {
		board[i] = make([]box, col)
	}
	return board
}

func CopyBoard(board multiAgentMatrix) multiAgentMatrix {
	row := len(board)
	col := len(board[0])
	board1 := InitializeBoard(row, col)

	for i := range board {
		for j := range board[0] {
			board1[i][j].IsAgent = board[i][j].IsAgent
			board1[i][j].IsFood = board[i][j].IsFood
			board1[i][j].foodChemo = board[i][j].foodChemo
			board1[i][j].trailChemo = board[i][j].trailChemo
			board1[i][j].agent = board[i][j].agent
		}
	}
	return board1
}
