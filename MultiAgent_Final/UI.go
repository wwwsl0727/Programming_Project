package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/therecipe/qt/widgets"
	"gocv.io/x/gocv"
)

type OrderPair struct {
	x, y int
}

func ImageShow(c chan image.Image, numGens int) {
	window := gocv.NewWindow("Simulation")
	count := 100
	defer window.Close()
	for true {
		img := <-c
		count += 100
		if count >= numGens+1 {
			break
		}

		fmt.Println(count)
		bounds := img.Bounds()
		x := bounds.Dx()
		y := bounds.Dy()

		bytes := make([]byte, 0, x*y)
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
				r, g, b, _ := img.At(i, j).RGBA()
				bytes = append(bytes, byte(b>>8))
				bytes = append(bytes, byte(g>>8))
				bytes = append(bytes, byte(r>>8))
			}
		}

		mat, err := gocv.NewMatFromBytes(y, x, gocv.MatTypeCV8UC3, bytes)
		if err != nil {
			panic("covert image to Mat error")
		}
		gocv.IMWrite("1.png", mat)
		time.Sleep(time.Second)
		im := gocv.IMRead("1.png", gocv.IMReadColor)

		window.IMShow(im)
		window.WaitKey(1)
	}
	<-c
}

func CreateNewImage(width, height int) {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	gray := color.RGBA{200, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, gray)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("test.png")
	png.Encode(f, img)

}

func AddFoodToImage() []OrderPair {
	dirName := "food"
	_, err := os.Stat(dirName)
	if err != nil {
		os.Mkdir(dirName, os.ModePerm)
	}

	fileWriter, err2 := os.Create(dirName + "/" + dirName + ".txt")
	if err2 != nil {
		panic("create food.txt error")
	}
	imageName := "test.png"
	winName := "food"
	size := image.Point{200, 200}

	foodList := make([]OrderPair, 0)
	img := gocv.IMRead(imageName, gocv.IMReadColor)
	gocv.Resize(img, &img, size, 0, 0, 0)
	window := gocv.NewWindow(winName)
	defer window.Close()
	radius := 5
	thickness := 10
	red := color.RGBA{255, 0, 0, 0}
	counter := 0

	var foodLocation OrderPair
	for true {
		rect := gocv.SelectROI(winName, img)
		counter++
		fmt.Println(rect.Min.X, rect.Min.Y)
		if rect.Min.X == 0 && rect.Min.Y == 0 && counter > 1 {
			break
		}
		if counter != 1 {
			foodLocation.x = rect.Min.X
			foodLocation.y = rect.Min.Y
			gocv.Circle(&img, rect.Min, radius, red, thickness)
			foodList = append(foodList, foodLocation)
		}
		window.WaitKey(1)
	}
	for _, food := range foodList {
		fmt.Fprintf(fileWriter, "%v %v", food.x, food.y)
		fmt.Fprintln(fileWriter)
	}
	gocv.IMWrite(imageName, img)
	return foodList
}

func AddWindToImage() []OrderPair {
	dirName := "wind"
	_, err := os.Stat(dirName)
	if err != nil {
		os.Mkdir(dirName, os.ModePerm)
	}

	fileWriter, err2 := os.Create(dirName + "/" + dirName + ".txt")
	if err2 != nil {
		panic("create wind.txt error")
	}
	imageName := "test.png"
	winName := "wind"
	windList := make([]OrderPair, 0)
	img := gocv.IMRead(imageName, gocv.IMReadColor)
	window := gocv.NewWindow(winName)
	defer window.Close()
	radius := 5
	thickness := 10
	red := color.RGBA{0, 100, 255, 0}
	counter := 0

	var windLocation OrderPair
	for true {
		rect := gocv.SelectROI(winName, img)
		counter++
		fmt.Println(rect.Min.X, rect.Min.Y)
		if rect.Min.X == 0 && rect.Min.Y == 0 && counter > 1 {
			break
		}
		if counter != 1 {
			windLocation.x = rect.Min.X
			windLocation.y = rect.Min.Y
			gocv.Circle(&img, rect.Min, radius, red, thickness)
			windList = append(windList, windLocation)
		}
		window.WaitKey(1)
	}
	for _, wind := range windList {
		fmt.Fprintf(fileWriter, "%v %v", wind.x, wind.y)
		fmt.Fprintln(fileWriter)
	}
	gocv.IMWrite(imageName, img)
	return windList
}

func AddLightToImage() []OrderPair {
	dirName := "light"
	_, err := os.Stat(dirName)
	if err != nil {
		os.Mkdir(dirName, os.ModePerm)
	}

	fileWriter, err2 := os.Create(dirName + "/" + dirName + ".txt")
	if err2 != nil {
		panic("create light.txt error")
	}
	imageName := "test.png"
	winName := "light"
	lightList := make([]OrderPair, 0)
	img := gocv.IMRead(imageName, gocv.IMReadColor)
	window := gocv.NewWindow(winName)
	defer window.Close()
	radius := 5
	thickness := 10
	red := color.RGBA{230, 230, 0, 0}
	counter := 0

	var lightLocation OrderPair
	for true {
		rect := gocv.SelectROI(winName, img)
		counter++
		fmt.Println(rect.Min.X, rect.Min.Y)
		if rect.Min.X == 0 && rect.Min.Y == 0 && counter > 1 {
			break
		}
		if counter != 1 {
			lightLocation.x = rect.Min.X
			lightLocation.y = rect.Min.Y
			gocv.Circle(&img, rect.Min, radius, red, thickness)
			lightList = append(lightList, lightLocation)
		}
		window.WaitKey(1)
	}
	for _, light := range lightList {
		fmt.Fprintf(fileWriter, "%v %v", light.x, light.y)
		fmt.Fprintln(fileWriter)
	}
	gocv.IMWrite(imageName, img)
	return lightList
}
func StartSimulation() {
	fmt.Println("start simmulation")
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

	numGens := 2000

	// emptyboard := InitializeBoard(row, col)
	matrix0 := InitializeBoard(row, col) // Used to pass to later simulation after initialization

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
	fmt.Println("All command line arguments read successfully.")

	c := make(chan image.Image, numGens/100)
	newboard := matrix0
	initImage := DrawGameBoard(newboard, 1, CN)
	c <- initImage
	go SimulateSlimeMold(matrix0, numGens, sensorArmLength, sensorAngle, sensorDiagonalL, depT, dampT, filterT, WL, WT, WN, CN, CL, dampN, filterN, RT, ET, c)
	ImageShow(c, numGens)
	CreateNewImage(300, 300)

}

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)
	foodList := make([]OrderPair, 0)
	windList := make([]OrderPair, 0)
	lightList := make([]OrderPair, 0)
	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Welcome to Slime Mold Simulation")
	window.SetMinimumSize2(300, 300)

	// Create main layout
	layout := widgets.NewQVBoxLayout()

	// Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create a line edit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	s := ""
	input.SetPlaceholderText("add some parameter location")
	layout.AddWidget(input, 0, 0)

	// Create a button and add it to the layout
	button1 := widgets.NewQPushButton2("click me to add food", nil)
	layout.AddWidget(button1, 0, 0)

	// Connect event for button
	button1.ConnectClicked(func(checked bool) {
		foodList = AddFoodToImage()
		s = "{"
		for _, foodLocation := range foodList {
			s += "(" + strconv.Itoa(foodLocation.x) + ", " + strconv.Itoa(foodLocation.y) + ")"
		}
		s += "}"
		input.SetPlaceholderText(s)
		widgets.QMessageBox_Information(nil, "OK", "press OK to finish selection",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

	})

	button2 := widgets.NewQPushButton2("click me to add wind", nil)
	layout.AddWidget(button2, 0, 0)

	button2.ConnectClicked(func(checked bool) {
		windList = AddWindToImage()
		s = "{"
		for _, windLocation := range windList {
			s += "(" + strconv.Itoa(windLocation.x) + ", " + strconv.Itoa(windLocation.y) + ")"
		}
		s += "}"
		input.SetPlaceholderText(s)
		widgets.QMessageBox_Information(nil, "OK", "press OK to finish selection",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

	})

	button3 := widgets.NewQPushButton2("click me to add light", nil)
	layout.AddWidget(button3, 0, 0)

	button3.ConnectClicked(func(checked bool) {
		lightList = AddLightToImage()
		s = "{"
		for _, lightLocation := range lightList {
			s += "(" + strconv.Itoa(lightLocation.x) + ", " + strconv.Itoa(lightLocation.y) + ")"
		}
		s += "}"
		input.SetPlaceholderText(s)
		widgets.QMessageBox_Information(nil, "OK", "press OK to finish selection",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

	})

	button4 := widgets.NewQPushButton2("click me to start simulation", nil)
	layout.AddWidget(button4, 0, 0)

	button4.ConnectClicked(func(checked bool) {
		StartSimulation()
	})

	CreateNewImage(300, 300)
	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)

	// Show the window
	window.Show()

	// Execute app
	app.Exec()

}
