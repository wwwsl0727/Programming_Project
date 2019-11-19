package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("MultiAgent Mold!")
	rand.Seed(time.Now().UnixNano())

	//intialization for 50% case
	row := 200
	col := 200
	sensorArmLength := 7
	sensorAngle:=math.Pi / float64(4)
	sensorDiagonalL := float64(5)*math.Sqrt(2)
	depT := 5    //The quantity of trail deposited by an agent
	dampT := 0.1 //Diffusion damping factor of trail
	// filter for trail 3x3

	//WT: The weight of trail value sensed by an agent’s sensor
	//WN: The weight of nutrient value sensed by an agent’s sensor
	WT := 0.4
	WN := 1 - WT

	CN := 10.0   //The Chemo-Nutrient concentration of food
	// dampN := 0.2 //Diffusion damping factor of chemo- nutrient
	// filter for nutrient 5x5

	//motion counter >= RT, reproduction
	//motion counter < RT, elimation
	RT := 15
	ET := 10

	//make the matrix which has 50% of its box has agent
	matrix := make(multiAgentMatrix, row)
	for i := range matrix {
		//for all boxes, intial foodChemo & trailChemo are 0.
		//IsFood, IsAgent are false
		row := make([]box, col)
		row = generateAgent(row, sensorArmLength,sensorDiagonalL,sensorAngle)
		matrix[i] = row
	}

	//put three food spots, each spot is a 3*3 matrix on the Matrix
	// matrix[10][100].IsFood = true
	// matrix[10][100].foodChemo = CN //10
	//
	// matrix[190][10].IsFood = true
	// matrix[190][10].foodChemo = CN //10
	//
	// matrix[190][190].IsFood = true
	// matrix[190][190].foodChemo = CN //10

	//The center is 10,100
	for i:=9;i<=11;i++{
		for j:=99;j<=101;j++{
			matrix[i][j].IsFood = true
			matrix[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,10
	for i:=189;i<=191;i++{
		for j:=9;j<=11;j++{
			matrix[i][j].IsFood = true
			matrix[i][j].foodChemo = CN //10
		}
	}

	//The center is 190,190
	for i:=189;i<=191;i++{
		for j:=189;j<=191;j++{
			matrix[i][j].IsFood = true
			matrix[i][j].foodChemo = CN //10
		}
	}
}

type multiAgentMatrix [][]box

type box struct {
	IsFood, IsAgent bool
	foodChemo       float64
	trailChemo      float64
	agent           *Agent
}

type Agent struct {
	motionCounter int // initial 0
	direction     float64 // direction could be the multiple of pi/4
	sensorLength  int // 7 in 50% agent case
	sensorDiagonalL float64 // float64(5)*math.Sqrt(2) for sensor in diagnol of axix
}

func generateAgent(row []box, sensorArmLength int,sensorDiagonalL,sensorAngle float64) []box {

	for i := range row {
		n := rand.Intn(2)                   //50% chance of having an agent in the box
		// randomDirection := rand.Intn(8) + 1 //This create a random direction from 1 - 8
		randomDirection := float64(rand.Intn(8) + 1) * sensorAngle //This create a random direction from pi/4 -2pi
		if n == 0 {
			row[i].IsAgent = true
			//create a new agent and points to it
			var oneAgent Agent
			oneAgent.motionCounter = 0
			oneAgent.direction = randomDirection
			oneAgent.sensorLength = sensorArmLength
			oneAgent.sensorDiagonalL = sensorDiagonalL
			row[i].agent = &oneAgent
		}
		// if n == 1, no agent
	}
	return row
}

//Compare two sensors measure, change a agent's direction
func (matrix multiAgentMatrix) SynthesisComparator(row, col int, WT, WN,sensorAngle float64) {
	//exceptions
	numRows := len(matrix)
	numCols := len(matrix[0])
	if row < 0 || col < 0 || row >= numRows || col >= numCols {
		panic("the given row and col is out of bound")
	}

	if matrix[row][col].IsAgent == false {
		panic("This box has no Agent!")
	}

	agentDirection:=matrix[row][col].agent.direction

	//For sensors on the x and y axis, sensorLength is 7
	L := matrix[row][col].agent.sensorLength

	//For sensors on the diagonal axis of xy plane, in order to get the location as integers, sensorlength is 5*sqart(2)
	DiagonalL := matrix[row][col].agent.sensorDiagonalL


	// The direction of the left,right sensor is based on the current direction of the agent.
	leftx,lefty,rightx,righty := CalculateSensorLocation(agentDirection, sensorAngle, DiagonalL, row, col, L)

	//Get sample chemoattractant values from sensors.
	// If left sensor and right sensor are all in the matrix
	if InField(numRows,numCols, leftx, lefty)  && InField(numRows,numCols, rightx, righty){
		sensorLeft := matrix[leftx][lefty]
		FL := calculateScore(sensorLeft, WT, WN)
		sensorRight := matrix[rightx][righty]
		FR := calculateScore(sensorRight, WT, WN)
		if FL<FR{
			//rotate right
			matrix[row][col].agent.direction -=sensorAngle
		}else if FL>FR{
				matrix[row][col].agent.direction +=sensorAngle
		}else{
			//random choose left and right
			if rand.Intn(2)==0{
					matrix[row][col].agent.direction +=sensorAngle
			}else{
					matrix[row][col].agent.direction -=sensorAngle
			}
		}
	}else if InField(numRows,numCols, leftx, lefty)== false && InField(numRows,numCols, rightx, righty){
		// If the left sensor is out of bound while the right sensor is in, turn right
		matrix[row][col].agent.direction -=sensorAngle
	}else if InField(numRows,numCols, leftx, lefty) && InField(numRows,numCols, rightx, righty)==false{
		//Otherwise, it turns left
			matrix[row][col].agent.direction +=sensorAngle
	}else{
		//If both side is out of bound, the agent turns back
		matrix[row][col].agent.direction += math.Pi
	}

}

//Calculate the location of left and right sensors.
 func CalculateSensorLocation(agentDirection,sensorAngle,DiagonalL float64,row,col,L int)(int,int,int,int){
	//Initialize the location of left, forward and right sensors
		 var leftx int
		 var lefty int
		 var rightx int
		 var righty int

	flag := int(agentDirection/sensorAngle)
	leftDirection:=agentDirection+ sensorAngle
	rightDirection:=agentDirection- sensorAngle

	//a is the coefficient for left and right arms; b is the coefficient for forward arms
	//a,b =L/DiagonalL
	var a float64
	//Check whether the forward direction of agent is +-pi/4;+-3pi/4;+-5pi/4;+-7pi/4;;
	//If so, a=L;leftx = row+a*cos(leftDirection); lefty= row+a*sin(leftDirection)
	//rightx = row+a*cos(rightDirection); righty= row+a*sin(rightDirection)
	//Else, exchange a=DiagonalL
	if flag % 2 !=0{
		 a=float64(L)
	}else{
		a=DiagonalL
	}
	leftx=row+int(a*math.Cos(leftDirection))
	lefty=col+int(a*math.Sin(leftDirection))
	rightx=row+int(a*math.Cos(rightDirection))
	righty=col+int(a*math.Sin(rightDirection))
	return leftx,lefty,rightx,righty
 }

//Calculate a box's score using SV = WT×TV +WN×NV
func calculateScore(B box, WT, WN float64) float64 {
	return WT*B.trailChemo + WN*B.foodChemo
}

//InField takes a the numRows and numCols of a matrix and i/j indices.  It returns true if (i,j) is a valid entry
//of the board.
func InField(numRows,numCols, i, j int) bool {
	if i < 0 || j < 0 || i >= numRows || j >= numCols {
		return false
	}
	// if we make it here, we are in the field.
	return true
}
