package main

import (
	"math"
	"math/rand"
)

func (board multiAgentMatrix) UpdateBoard(boardCopy multiAgentMatrix,
	sensorArmLength int,
	sensorAngle float64,
	sensorDiagonalL float64,
	depT float64,
	dampT float64,
	filterN int,
	WL float64,
	WT float64,
	WN float64,
	CN float64,
	CL float64,
	dampN float64,
	filterT int,
	RT int,
	ET int,
	rowIndices []int,
	colIndices []int) {

	//For boundary,both the foodChemo and the trailChemo is 0. Use filters to simulate the process of chemical spread
	//among the 2D board.
	board.UpdateChemo(filterN, dampN, boardCopy, "food")

	board.UpdateChemo(filterT, dampT, boardCopy, "trial")

	//The previous move the left and up corner is because the agents is updated from left to right and
	//from above to bottom. Therefore, now we update each agent in a random order. Also in order not to get the
	//boundary, the x and y shuffle from [1,2,...len(board)-2] instead of [0,1,...len(board-1)].
	rand.Shuffle(len(rowIndices), func(i, j int) {
		rowIndices[i], rowIndices[j] = rowIndices[j], rowIndices[i]
	})
	rand.Shuffle(len(colIndices), func(i, j int) {
		colIndices[i], colIndices[j] = colIndices[j], colIndices[i]
	})

	//update agent
	for i := range rowIndices {
		for j := range colIndices {
			//sense direction and change direction to the higher chemo direction.
			r := rowIndices[i]
			c := colIndices[j]
			if boardCopy[r][c].IsAgent {
				agentDirection := board.SynthesisComparator(r, c, WT, WN, WL, sensorAngle)

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

				forwardx = r + int(math.Round(a*math.Cos(agentDirection)))
				forwardy = c + int(math.Round(a*math.Sin(agentDirection)))
				currCell := &board[r][c]
				forwardcell := &board[forwardx][forwardy]

				//If the forward direction is occupied, change direction randomly and motionCounter--
				if forwardcell.IsAgent == true {
					currCell.agent.direction = float64(rand.Intn(8)+1) * sensorAngle
					currCell.agent.motionCounter--
					if currCell.agent.motionCounter < ET {
						//currentCell die
						currCell.IsAgent = false
					}
				} else {
					//If the forward direction is not occupied, move to that direction and leave trail in the new cell
					forwardcell.IsAgent = true
					forwardcell.trailChemo += depT
					var forwardAgent Agent
					forwardAgent.direction = currCell.agent.direction
					forwardAgent.motionCounter = currCell.agent.motionCounter + 1
					//rotate to the direction with the higher sense value
					forwardAgent.sensorDiagonalL = sensorDiagonalL
					forwardAgent.sensorLength = sensorArmLength
					forwardcell.agent = forwardAgent

					forwardcell.agent.direction = board.SynthesisComparator(forwardx, forwardy, WT, WN, WL, sensorAngle)
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

}

//Update foodchemo/trial chemo in the 2d board by filters and damping.
func (board multiAgentMatrix) UpdateChemo(filter int, damp float64, boardCopy multiAgentMatrix, category string) {
	//use 5*5 average filter, new board is updated based on the previous board
	board.AverageFilter(filter, category, boardCopy)
	//damp
	board.Damp(damp, category)

}

//Update foodchemo/trial chemo in the 2d board by filters.
func (board multiAgentMatrix) AverageFilter(filter int, category string, boardCopy multiAgentMatrix) {
	numRows := len(board)
	numCols := len(board[0])
	for i := range board {
		for j := range board[0] {
			//check the neighbors and get the Average.
			// For boundary,both the foodChemo and the trailChemo is 0.
			//Apply food filters only on box which is not food resource
			if i != 0 && i != numRows-1 && j != 0 && j != numCols-1 {
				if category == "food" {
					if !board[i][j].IsFood {
						board.AverageNeighbor(i, j, filter, category, boardCopy)
					}
				} else if category == "trial" {
					board.AverageNeighbor(i, j, filter, category, boardCopy)
				} else {
					panic("the given category is wrong")
				}
			}
		}
	}
}

//Update the food/trial chemo based on the chemicals of its neighbors.
func (board multiAgentMatrix) AverageNeighbor(r, c, filter int, category string, boardCopy multiAgentMatrix) {
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

	}
}

//The trial/food chemo degenerates as time goes by.
func (board multiAgentMatrix) Damp(damp float64, category string) {
	factor := 1.0 - damp
	for i := range board {
		for j := range board[i] {
			//trial value of food source can be changed
			if category == "food" {
				if !board[i][j].IsFood {

					board[i][j].foodChemo *= factor
				}
			} else {
				board[i][j].trailChemo *= factor
			}
		}
	}
}

//GenerateAgent Generate agent for a row. 50% of boxes have agent
func GenerateAgent(row []box, sensorArmLength int, sensorDiagonalL, sensorAngle float64) []box {

	for i := 1; i < len(row)-1; i++ {
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

	agentDirection := matrix[row][col].agent.direction

	//For sensors on the x and y axis, sensorLength is 7
	L := matrix[row][col].agent.sensorLength

	//For sensors on the diagonal axis of xy plane, in order to get the location as integers, sensorlength is 5*sqart(2)
	DiagonalL := matrix[row][col].agent.sensorDiagonalL

	// The direction of the left,right sensor is based on the current direction of the agent.
	leftx, lefty, rightx, righty := CalculateSensorLocation(agentDirection, sensorAngle, DiagonalL, row, col, L)

	//Get sample chemoattractant values from sensors.
	// If left sensor and right sensor are all in the matrix
	var FL float64
	var FR float64
	if InField(numRows, numCols, leftx, lefty) && InField(numRows, numCols, rightx, righty) {
		sensorLeft := matrix[leftx][lefty]
		FL = calculateScore(sensorLeft, WT, WN, WL)
		sensorRight := matrix[rightx][righty]
		FR = calculateScore(sensorRight, WT, WN, WL)
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
	// fmt.Printf("left sensor with location %d %d, chemo is %f\n", leftx, lefty, FL)
	// fmt.Printf("right sensor with location %d %d, chemo is %f\n", rightx, righty, FR)
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
	leftx = row + int(math.Round(a*math.Cos(leftDirection)))
	lefty = col + int(math.Round(a*math.Sin(leftDirection)))
	rightx = row + int(math.Round(a*math.Cos(rightDirection)))
	righty = col + int(math.Round(a*math.Sin(rightDirection)))
	return leftx, lefty, rightx, righty
}

//Calculate a box's score using SV = WT×TV +WN×NV
func calculateScore(B box, WT, WN, WL float64) float64 {
	return WT*B.trailChemo + WN*B.foodChemo - WL*B.light
}

//InField takes a the numRows and numCols of a matrix and i/j indices.  It returns true if (i,j) is a valid entry
//of the board.
func InField(numRows, numCols, i, j int) bool {
	if i <= 0 || j <= 0 || i >= numRows-1 || j >= numCols-1 {
		return false
	}
	// if we make it here, we are in the field.
	return true
}

//Initialize the board as a row*col multiAgentMatrix.
func InitializeBoard(row, col int) multiAgentMatrix {
	board := make(multiAgentMatrix, row)
	for i := range board {
		board[i] = make([]box, col)
	}
	return board
}

//Copy the board and return the new board with the same fields and values but different address.
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
			board1[i][j].light = board[i][j].light
			board1[i][j].haslight = board[i][j].haslight

		}
	}
	return board1
}
