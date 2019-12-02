package main

import (
	"math"
	"math/rand"
)

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
	//The center is 1,1
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 198,1
	for i := 197; i <= 199; i++ {
		for j := 0; j <= 2; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 198,198
	for i := 197; i <= 197; i++ {
		for j := 197; j <= 197; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}
	//The center is 1,198
	for i := 0; i <= 2; i++ {
		for j := 197; j <= 199; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}
	// //The center is 100,50
	// for i := 99; i <= 101; i++ {
	// 	for j := 49; j <= 51; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //The center is 40,150
	// for i := 39; i <= 41; i++ {
	// 	for j := 149; j <= 151; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //The center is 160,150
	// for i := 159; i <= 161; i++ {
	// 	for j := 149; j <= 151; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //food in light
	// for i := 99; i <= 101; i++ {
	// 	for j := 99; j <= 101; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }

	//Given the center of light, the boxes in light source has stronger light concentration
	addlight(matrix0, x, y, CL)

	return matrix0
}

//Add light concentration to a square of length 9
func addlight(matrix0 multiAgentMatrix, x, y int, CL float64) {
	for i := x - 40; i <= x+40; i++ {
		for j := y - 40; j <= y+40; j++ {
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
	/*
		//The center is 100,50
		for i := 99; i <= 101; i++ {
			for j := 49; j <= 51; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}

		//The center is 40,150
		for i := 39; i <= 41; i++ {
			for j := 149; j <= 151; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}

		//The center is 160,150
		for i := 159; i <= 161; i++ {
			for j := 149; j <= 151; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}
	*/
	//The center is 100,150
	// for i := 99; i <= 101; i++ {
	// 	for j := 149; j <= 151; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	// //The center is 160,50
	// for i := 159; i <= 161; i++ {
	// 	for j := 49; j <= 51; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }
	//
	// //The center is 40,50
	// for i := 39; i <= 41; i++ {
	// 	for j := 49; j <= 51; j++ {
	// 		matrix0[i][j].IsFood = true
	// 		matrix0[i][j].foodChemo = CN //10
	// 	}
	// }

	//The center is 1,1
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 198,1
	for i := 197; i <= 199; i++ {
		for j := 0; j <= 2; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 198,198
	for i := 197; i <= 197; i++ {
		for j := 197; j <= 197; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}
	//The center is 1,198
	for i := 0; i <= 2; i++ {
		for j := 197; j <= 199; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	return matrix0
}

func intializeCornerBoard(matrix0 multiAgentMatrix, row, col, sensorArmLength int, sensorDiagonalL, sensorAngle, CN float64) multiAgentMatrix {

	//The center is 36,51
	for i := 35; i <= 37; i++ {
		for j := 50; j <= 52; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 85,80
	for i := 84; i <= 86; i++ {
		for j := 79; j <= 81; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 140,60
	for i := 139; i <= 141; i++ {
		for j := 59; j <= 61; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 70,130
	for i := 69; i <= 71; i++ {
		for j := 129; j <= 131; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 130,146
	for i := 129; i <= 131; i++ {
		for j := 145; j <= 147; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 179,119
	for i := 178; i <= 181; i++ {
		for j := 118; j <= 120; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}

	//The center is 99,184
	for i := 98; i <= 100; i++ {
		for j := 183; j <= 185; j++ {
			matrix0[i][j].IsFood = true
			matrix0[i][j].foodChemo = CN //10
		}
	}
	/*
		//The center is 1,1
		for i := 0; i <= 2; i++ {
			for j := 0; j <= 2; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}

		//The center is 198,1
		for i := 197; i <= 199; i++ {
			for j := 0; j <= 2; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}

		//The center is 198,198
		for i := 197; i <= 197; i++ {
			for j := 197; j <= 197; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}
		//The center is 1,198
		for i := 0; i <= 2; i++ {
			for j := 197; j <= 199; j++ {
				matrix0[i][j].IsFood = true
				matrix0[i][j].foodChemo = CN //10
			}
		}
	*/
	//	put 9 agents in the southern part,100,190
	for i := 99; i <= 101; i++ {
		for j := 189; j <= 191; j++ {
			matrix0[i][j].IsAgent = true
			var agent Agent
			//flag := float64(rand.Intn(8) + 1)
			//flag := 3.0
			//agent.direction = flag * sensorAngle
			//fmt.Println("initial direction", flag)
			agent.direction = float64(rand.Intn(8)+1) * (math.Pi / float64(4))
			agent.motionCounter = 0
			agent.sensorDiagonalL = sensorDiagonalL
			agent.sensorLength = sensorArmLength
			matrix0[i][j].agent = agent
		}
	}

	return matrix0
}
