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
	SensorArmLength := 7
	depT := 5    //The quantity of trail deposited by an agent
	dampT := 0.1 //Diffusion damping factor of trail
	// filter for trail 3x3

	//WT: The weight of trail value sensed by an agent’s sensor
	//WN: The weight of nutrient value sensed by an agent’s sensor
	WT := 0.4
	WN := 1 - WT

	CN := 10     //The Chemo-Nutrient concentration of food
	dampN := 0.2 //Diffusion damping factor of chemo- nutrient
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
		row = generateAgent(row, SensorArmLength)
		matrix[i] = row
	}
	//put three food spots on the Matrix
	matrix[10][100].IsFood = true
	matrix[10][100].foodChemo = 10

	matrix[190][10].IsFood = true
	matrix[190][10].foodChemo = 10

	matrix[190][190].IsFood = true
	matrix[190][190].foodChemo = 10
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
	direction     int // direction could be 1 - 8, represent 8 direction
	sensorLength  int // 7 in 50% agent case
}

func generateAgent(row []box, SensorArmLength int) []box {

	for i := range row {
		n := rand.Intn(2)                   //50% chance of having an agent in the box
		randomDirection := rand.Intn(8) + 1 //This create a random direction from 1 - 8
		if n == 0 {
			row[i].IsAgent = true
			//create a new agent and points to it
			var oneAgent Agent
			oneAgent.motionCounter = 0
			oneAgent.direction = randomDirection
			oneAgent.sensorLength = SensorArmLength
			row[i].agent = &oneAgent
		}
		// if n == 1, no agent
	}
	return row
}

//Compare three sensors measure, change a agent's direction
func (matrix multiAgentMatrix) SynthesisComparator(row, col int, WT, WN float64) {
	if matrix[row][col].IsAgent == false {
		panic("This box has no Agent!")
	}

	sensorLength := matrix[row][col].agent.sensorLength

	//three sensors location
	sensor1 := matrix[row-sensorLength][col-sensorLength]
	sensor2 := matrix[row-sensorLength][col]
	sensor3 := matrix[row-sensorLength][col+sensorLength]

	sensor1Score := calculateScore(sensor1, WT, WN)
	sensor2Score := calculateScore(sensor2, WT, WN)
	sensor3Score := calculateScore(sensor3, WT, WN)

	//Find the biggest of three scores and change agent's direction
	max := FindMax(sensor1Score, sensor2Score, sensor3Score)
	if max == 1 {
		matrix[row][col].agent.direction -= 1
	}
	// if max == 2 {
	// 	matrix[row][col].agent.direction
	// }
	if max == 3 {
		matrix[row][col].agent.direction += 1
	}
}

//Calculate a box's score using SV = WT×TV +WN×NV
func calculateScore(B box, WT, WN float64) float64 {
	return WT*B.trailChemo + WN*B.foodChemo
}

func FindMax(score1, score2, score3 float64) int {
	var slice = []float64{score1, score2, score3}
	max := math.Max(score1, math.Max(score2, score3))

	for i := range slice {
		if slice[i] == max {
			return i
		}
	}
	panic("Error when finding Max value")
}
