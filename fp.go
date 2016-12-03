package main

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// GameBoard: The game board is a 2D slice of Cell objects//
type GameBoard [][]float64

// TuringPattern represent different Turing Scale
type TuringPattern struct {
	activator                GameBoard
	inhibitor                GameBoard
	variations               GameBoard
	activatorRX, activatorRY int
	inhibitorRX, inhibitorRY int
	variationSampleR         int
	stepSize                 float64
	num                      int
	tmpDestA                 GameBoard
	tmpDestB                 GameBoard
	blurSteps                int
	updateType               int // 0 - fast, 1 - rectangle
}

//initializeTuringPattern() method will initialize the TuringPattern of certain scale
func (tp *TuringPattern) initializeTuringPattern(num, aRX, aRY, iRX, iRY, vSR,
	blurType, blurSteps int, stepSize float64) {
	tp.num = num
	tp.activator = make2dSliceOfFloatNum(num)
	tp.inhibitor = make2dSliceOfFloatNum(num)
	tp.tmpDestA = make2dSliceOfFloatNum(num)
	tp.tmpDestB = make2dSliceOfFloatNum(num)
	tp.variations = make2dSliceOfFloatNum(num)
	tp.activatorRX = aRX
	tp.activatorRY = aRY
	tp.inhibitorRX = iRX
	tp.inhibitorRY = iRY
	tp.variationSampleR = vSR
	tp.stepSize = stepSize
	tp.updateType = blurType
	tp.blurSteps = blurSteps
}

// make a 2d slice of float64
func make2dSliceOfFloatNum(num int) [][]float64 {
	twoDSlice := make(GameBoard, num)
	for i := 0; i < num; i++ {
		oneDSlice := make([]float64, num)
		twoDSlice[i] = oneDSlice
	}
	return twoDSlice
}

// initialize different scale turing patterns
func initializePatterns(patterns []TuringPattern, num, blurType, blurSteps, numOfPatterns int) []TuringPattern {
	var pattern4 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern4.initializeTuringPattern(num, 100, 100,
		200, 200, 1, blurType, blurSteps, 0.5)

	//pattern4.initializeTuringPattern(num, 70, 70,
	//	200, 200, 1, blurType, blurSteps, 0.5)

	var pattern3 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern3.initializeTuringPattern(num, 20, 20,
		40, 40, 1, blurType, blurSteps, 0.4)
	//pattern3.initializeTuringPattern(num, 20, 20,
	//	60, 60, 1, blurType, blurSteps, 0.4)

	var pattern2 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern2.initializeTuringPattern(num, 10, 10,
		20, 20, 1, blurType, blurSteps, 0.3)
	//pattern2.initializeTuringPattern(num, 10, 10,
	//	30, 30, 1, blurType, blurSteps, 0.3)

	var pattern1 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern1.initializeTuringPattern(num, 5, 5,
		10, 10, 1, blurType, blurSteps, 0.2)
	//pattern1.initializeTuringPattern(num, 5, 5,
	//15, 15, 1, blurType, blurSteps, 0.2)

	var pattern0 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern0.initializeTuringPattern(num, 1, 1,
		2, 2, 1, blurType, blurSteps, 0.1)
	//pattern0.initializeTuringPattern(num, 1, 1,
	//	3, 3, 1, blurType, blurSteps, 0.1)
	for i := 0; i < numOfPatterns; i++ {
		switch i {
		case 0:
			patterns = append(patterns, pattern0)
		case 1:
			patterns = append(patterns, pattern1)
		case 2:
			patterns = append(patterns, pattern2)
		case 3:
			patterns = append(patterns, pattern3)
		case 4:
			patterns = append(patterns, pattern4)
		}
	}
	return patterns
}

// create gameboard type board, with size boardSize*boardSize
func createGameBoard(boardSize int) GameBoard {
	var board GameBoard
	// need to create each board row
	for i := 0; i < boardSize; i++ {
		boardRow := make([]float64, boardSize)
		board = append(board, boardRow)
	}
	return board
}

// initializeBoard with random float number in range(-1, 1)
func initializeBoard(board GameBoard) GameBoard {
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			board[i][j] = rand.Float64()*2 - 1
		}
	}
	//fmt.Println("after initialize=", board)
	return board
}

// update blur steps and levels
// basically what it does is:
//  blurSteps=max(1,blurSteps)
//  levels = min(5,max(1,levels))
// !!!! need to work on this
func updateBlurStepsAndLevels(blurSteps, levels int) (int, int) {
	if blurSteps < 1 {
		blurSteps = 1
	}
	if levels < 1 {
		levels = 1
	}
	if levels > 5 {
		levels = 5
	}
	return blurSteps, levels
}

// readParameters() get input fron terminal
func readParameters() (int, int, int, int, int) {
	// get parameters
	parameters := os.Args
	if parameters == nil || len(parameters) != 6 {
		fmt.Println("Error: 5 parameters are wanted.")
		fmt.Println("Please Input 5 number to represent: ")
		fmt.Println("1) imageSize(ie, 256)")
		fmt.Println("2) updateType(ie, 0->fast; 1->rectangle; 2->circle)")
		fmt.Println("3) step Number(ex. 100)")
		fmt.Println("4) turing Scale Num(less than 5)")
		fmt.Println("5) draw period")

		fmt.Println("Please refer readMe.txt if have more question.")
		os.Exit(1)
	}
	// parameter convert
	boardSize, err1 := strconv.Atoi(parameters[1])
	blurType, err2 := strconv.Atoi(parameters[2])
	stepNum, err3 := strconv.Atoi(parameters[3])
	scaleNum, err4 := strconv.Atoi(parameters[4])
	drawPeriod, err5 := strconv.Atoi(parameters[5])
	// see if there is any illegal input
	if err1 != nil || boardSize < 0 {
		fmt.Println("Error: something wrong with image size you input.")
		fmt.Println("please input an interger within in range(256, 1024)")
		os.Exit(1)
	}
	if err2 != nil || blurType < 0 || blurType > 2 {
		fmt.Println("Error: something wrong with blurType you input.")
		fmt.Println("The blurType can be 0, 1 or 2")
		fmt.Println("Type(ie, 0->fast; 1->rectangle; 2->circle); 3)step Number(ex. 100); 4)turing Scale Num(less than 5)")
		os.Exit(1)
	}
	if err3 != nil || stepNum < 0 {
		fmt.Println("Error: something wrong with the step number.")
		fmt.Println("Input a number in range of (20, 100)")
		os.Exit(1)
	}
	if err4 != nil || scaleNum < 0 {
		fmt.Println("Error: something wrong with the scale number")
		fmt.Println("input a number in range of (20, 100)")
		os.Exit(1)
	}
	if err5 != nil || scaleNum < 0 {
		fmt.Println("Error: something wrong with the drawperiod")
		fmt.Println("input a positive integer to represent the draw period you want")
		os.Exit(1)
	}
	return boardSize, blurType, stepNum, scaleNum, drawPeriod
}

// update grid board[row][col] from the patterns[i].variations[row][col]
// 1. Find which of the scales has the smallest variation value.
//ie find which scale has the lowest variation[x,y,scalenum] value and call this bestvariation
//2. Using the scale with the smallest variation, update the grid value
//if activator[row][col][bestvariationscale]>inhibitor[row][col][bestvariationscale]>
//then grid[row][col]:=grid[row][co]+smallamounts[bestvariation]
//else grid[row][co]:=grid[row][co]-smallamounts[bestvariation]
//
func updateBoardFromPatterns(patterns []TuringPattern, board GameBoard) {
	rows := len(board)
	cols := len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] += bestVariationFromPatterns(row, col, patterns)
		}
	}
}

//if activator[row][col][bestvariationscale]>inhibitor[row][col][bestvariationscale]>
//then return: +smallamounts[bestvariation]； else return: -smallamounts[bestvariation]
//the return value should include sign(positive negative)
func bestVariationFromPatterns(row, col int, patterns []TuringPattern) float64 {
	scaleNums := len(patterns)
	// assume the patterns[0] is the best turing scale
	bestScale := 0
	sign := 1.0
	inhibitor := patterns[bestScale].inhibitor[row][col]
	activator := patterns[bestScale].activator[row][col]
	if inhibitor < activator {
		sign = -1.0
	} else {
		sign = 1.0
	}
	bestVariation := patterns[bestScale].variations[row][col]
	// loop throught all scale to find the actual best and return
	for i := 0; i < scaleNums; i++ {
		if patterns[i].variations[row][col] < bestVariation {
			bestVariation = patterns[i].variations[row][col]
			bestScale = i
			inhibitor = patterns[bestScale].inhibitor[row][col]
			activator = patterns[bestScale].activator[row][col]
			if inhibitor > activator {
				sign = -1.0
			} else {
				sign = 1.0
			}
		}
	}
	return bestVariation * sign
}

// use patterns[i].activator[][] and inhibitor[i].inhibitor[][] to
//calculate variations variation=variation+abs(activator[x][y]-inhibitor[x][y])
// patterns[i].variationSampleR should be used as the range of sample
// patterns[i].variationSampleR = 1 by default
func updateScalesVariation(patterns []TuringPattern, board GameBoard) {
	// loop through all the scales of turing pattern and calculate
	//variation=variation+abs(activator[x][y]-inhibitor[x][y])
	scaleNum := len(patterns)
	rows := len(patterns[0].variations)
	cols := len(patterns[0].variations[0])
	for i := 0; i < scaleNum; i++ {
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				patterns[i].variations[row][col] +=
					math.Abs(patterns[i].activator[row][col] -
						patterns[i].inhibitor[row][col])
			}
		}
	}
}

// will enter different update mode from here, based on the blurType
func updateTuringScales(patterns []TuringPattern, board GameBoard) {
	// splash different update mode:
	//0--> fast mode;
	//1--> rectangle mode;
	//2-->circle mode(slow)
	switch patterns[0].updateType {
	case 0:
		// update based on the row and col of the cell
		quickUpdateScales(patterns, board)
	case 1:
		// update based on (row-r, col-r, row+r, col +r)
		// is more accurate and more time consuming
		rectUpdateScales(patterns, board)
	case 2:
		circleUpdateScales(patterns, board)
		//gussianUpdateScales(patterns, board)
	}
}

///////////////////////////////////////////////////////////////////////
// circle update mode, the most intuitive, but also a little slower method
// use the (row-R, col-R, cow+R, col+R) as the outer range
// of calculate ingredient[row][col]
///////////////////////////////////////////////////////////////////////
func circleUpdateScales(patterns []TuringPattern, board GameBoard) {
	scaleNum := len(patterns)
	for i := 0; i < scaleNum; i++ {
		patterns[i].circleUpdateScale(board)
	}
}

// circleUpdateScale
// udpate the board[row][col] from the value (row-R, col-R, cow+R, col+R)
func (turingScale *TuringPattern) circleUpdateScale(board GameBoard) {
	//updateActivator, row and col
	circleUpdateActivator(turingScale.activator, turingScale.activatorRX, board)
	//update Inhibitor, row and col
	circleUpdateInhibitor(turingScale.inhibitor, turingScale.inhibitorRX, board)
}

func circleUpdateActivator(ingredient [][]float64, radius int, board GameBoard) {
	rows := len(board)
	cols := len(board[0])
	//fmt.Println("here!", rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			count := 0
			// use the (row-R, col-R, cow+R, col+R) as the outer range
			// of calculate ingredient[row][col]
			for i := row - radius; i < row+radius; i++ {
				if i >= 0 && i < rows {
					for j := col - radius; j < col+radius; j++ {
						if j >= 0 && j < cols {
							// use the specific range to find if the cell is
							// in circiel (row, col) with Radius

							if ((i-row)*(i-row) + (j-col)*(j-col)) < radius*radius {
								// for wrap around
								currRow := (i + rows) % rows
								currCol := (j + cols) % cols
								for currRow < 0 {
									currRow += rows
								}
								for currCol < 0 {
									currCol += cols
								}
								//fmt.Println("radius=", radius)
								//fmt.Println("i, j=", i, j)
								//fmt.Println("currRow currCol=", currRow, currCol)
								//fmt.Println("Rows Cols=", rows, cols)

								sum += board[i][j]
								count++
							}
						}
					}
				}
			}
			ingredient[row][col] = sum / float64(count)
			//fmt.Println("ingredient=", ingredient[row][col])
		}
	}
}

func circleUpdateInhibitor(ingredient [][]float64, radius int, board GameBoard) {
	rows := len(board)
	cols := len(board[0])
	//fmt.Println("here!", rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			count := 0

			// use the (row-R, col-R, cow+R, col+R) as the outer range
			// of calculate ingredient[row][col]
			for i := row - radius; i < row+radius; i++ {
				if i >= 0 && i < rows {
					for j := col - radius; j < col+radius; j++ {
						if j >= 0 && j < cols {
							// use the specific range to find if the cell is
							// in circiel (row, col) with Radius

							if ((i-row)*(i-row) + (j-col)*(j-col)) < radius*radius {
								// for wrap around
								currRow := (i + rows) % rows
								currCol := (j + cols) % cols
								for currRow < 0 {
									currRow += rows
								}
								for currCol < 0 {
									currCol += cols
								}
								//fmt.Println("radius=", radius)
								//fmt.Println("i, j=", i, j)
								//fmt.Println("currRow currCol=", currRow, currCol)
								//fmt.Println("Rows Cols=", rows, cols)

								sum += board[i][j]
								count++
							}
						}
					}
				}
			}
			ingredient[row][col] = sum / float64(count)
			//fmt.Println("ingredient=", ingredient[row][col])
		}
	}
}

// check if (i, j) is legal cell coordinate on board
func onBoard(i, j int, board GameBoard) bool {
	rows := len(board)
	cols := len(board[0])
	if (i >= 0) && (j >= 0) && (i < rows) && (j < cols) {
		return true
	}
	return false
}

///////////////////////////////////////////////////////////////////////
// rectangle update mode
// use the (row-R, col-R, cow+R, col+R) as the scope of calculate ingredient[row][col]
// could be slower
///////////////////////////////////////////////////////////////////////
func rectUpdateScales(patterns []TuringPattern, board GameBoard) {
	scaleNum := len(patterns)
	for i := 0; i < scaleNum; i++ {
		patterns[i].rectUpdateScale(board)
	}
}

func (turingScale *TuringPattern) rectUpdateScale(board GameBoard) {
	//updateActivator, row and col
	rectUpdateActivator(turingScale.activator, board, turingScale.activatorRX)
	//update Inhibitor, row and col
	rectUpdateInhibitor(turingScale.inhibitor, board, turingScale.inhibitorRX)
}

// udpate activator[][] in recttnagle way
func rectUpdateActivator(ingredient, board GameBoard, radius int) {
	// get board dimenstions
	rows := len(ingredient)
	cols := len(ingredient[0])
	// use the (row-R, col-R, cow+R, col+R) as the scope of calculate ingredient[row][col]
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			count := 0
			for i := row - radius; i < row+radius; i++ {
				for j := col - radius; j < col+radius; j++ {
					// no wrap around for this mode
					if i >= 0 && j >= 0 && i < rows && j < cols {
						currRow := (i + rows) % rows
						currCol := (j + cols) % cols
						sum += board[currRow][currCol]
						count++
					}
				}
				ingredient[row][col] = sum / float64(count)
			}
		}
	}
}

// udpate inhibitor[][] in recttnagle way
func rectUpdateInhibitor(ingredient, board GameBoard, radius int) {
	// get board dimenstions
	rows := len(ingredient)
	cols := len(ingredient[0])
	// use the (row-R, col-R, cow+R, col+R) as the scope of calculate ingredient[row][col]
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			count := 0
			for i := row - radius; i < row+radius; i++ {
				for j := col - radius; j < col+radius; j++ {
					currRow := (i + rows) % rows
					currCol := (j + cols) % cols
					sum += board[currRow][currCol]
					count++
				}
			}
			ingredient[row][col] = sum / float64(count)
		}
	}
}

///////////////////////////////////////////////////////////////////////
// quick update mode
// only use the the cell in range (row-R, row+r) and (col-R, col+R)
///////////////////////////////////////////////////////////////////////
func quickUpdateScales(patterns []TuringPattern, board GameBoard) {
	scaleNum := len(patterns)
	for i := 0; i < scaleNum; i++ {
		patterns[i].quickUpdateScale(board)
	}
}

// only use the the cell in range (row-R, row+r) and (col-R, col+R)
func (turingScale *TuringPattern) quickUpdateScale(board GameBoard) {
	//updateActivator, row and col
	turingScale.updateActivatorRow(board)
	turingScale.updateActivatorCol(board)
	//update Inhibitor, row and col
	turingScale.updateInhibitorRow(board)
	turingScale.updateInhibitorCol(board)
}

// update inhibitor[][]-----Row
func (turingScale *TuringPattern) updateInhibitorRow(board GameBoard) {
	rows := len(turingScale.inhibitor)
	cols := len(turingScale.inhibitor[0])
	actR := turingScale.inhibitorRY
	// calculate sum
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for i := row - actR; i < row+actR; i++ {
				// wrap around
				currRow := (i + rows)
				for currRow < 0 {
					currRow += rows
				}
				currRow = currRow % rows
				sum += board[currRow][col]
			}
			turingScale.inhibitor[row][col] = sum / float64(actR*2+1)
		}
	}
}

// update inhibitor[][]-----col
func (turingScale *TuringPattern) updateInhibitorCol(board GameBoard) {
	rows := len(turingScale.inhibitor)
	cols := len(turingScale.inhibitor[0])
	actR := turingScale.inhibitorRX
	// calculate sum
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for j := col - actR; j < col+actR; j++ {
				// wrap around
				currCol := j + cols
				for currCol < 0 {
					currCol += cols
				}
				currCol = currCol % cols
				sum += board[row][currCol]
			}
			turingScale.inhibitor[row][col] =
				(turingScale.inhibitor[row][col] + sum/float64(actR*2+1)) / 2.0
		}
	}
}

// update activator[][]-----> Row
func (turingScale *TuringPattern) updateActivatorRow(board GameBoard) {
	rows := len(turingScale.activator)
	cols := len(turingScale.activator[0])
	actR := turingScale.activatorRY
	// calculate sum
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for i := row - actR; i < row+actR; i++ {
				// wrap around
				currRow := i + rows
				for currRow < 0 {
					currRow += rows
				}
				currRow = currRow % rows
				sum += board[currRow][col]
			}
			turingScale.activator[row][col] = sum / float64(actR*2+1)
		}
	}
}

// update activator[][]-----> Col
func (turingScale *TuringPattern) updateActivatorCol(board GameBoard) {
	rows := len(turingScale.activator)
	cols := len(turingScale.activator[0])
	actR := turingScale.activatorRX
	// calculate sum
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for j := col - actR; j < col+actR; j++ {
				// wrap around
				currCol := j + cols
				for currCol < 0 {
					currCol += cols
				}
				currCol = currCol % cols
				sum += board[row][currCol]
			}
			turingScale.activator[row][col] =
				(turingScale.activator[row][col] + sum/float64(actR*2+1)) / 2.0
		}
	}
}

// normalizeBoard() scale the value on baord bacj to [-1, 1]
func normalizeBoard(board GameBoard) {
	rows := len(board)
	cols := len(board[0])
	// find the minValue and maxValue on board
	minValue := minOfBoard(board)
	maxValue := maxOfBoard(board)

	units := (maxValue - minValue) / 2.0
	// compute the noemalized value
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			value := board[row][col]
			value = (value-minValue)/units - 1.0
			board[row][col] = value
		}
	}
}

// find the max value on board
func maxOfBoard(board GameBoard) float64 {
	rows := len(board)
	cols := len(board[0])
	//maxValue on board
	maxValue := 0.0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] > maxValue {
				maxValue = board[row][col]
			}
		}
	}
	return maxValue
}

// find the min value on board
func minOfBoard(board GameBoard) float64 {
	rows := len(board)
	cols := len(board[0])
	//maxValue on board
	minValue := 0.0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] < minValue {
				minValue = board[row][col]
			}
		}
	}
	return minValue
}

func main() {
	// set up many variables
	// blurType(update method) and stepNum and boardSize should be get from command line input
	boardSize, blurType, stepNum, scaleNum, drawPeriod := readParameters()
	num := boardSize
	levels := 1
	blurSteps := 5
	//maxCount := 10000
	//convergeThreshold := 0.01 //1
	blurSteps, levels = updateBlurStepsAndLevels(blurSteps, levels)
	// patterns is a slice of different turing patterns
	patterns := make([]TuringPattern, 0)
	patterns = initializePatterns(patterns, num, blurType, blurSteps, scaleNum)
	// create the grid, with each [][]float64
	turingBoard := createGameBoard(boardSize)
	// fill the board with random float64 in range(-1,1)
	turingBoard = initializeBoard(turingBoard)
	//diffBoardSum := 1
	calculateTuringPatternBoard(patterns, turingBoard, stepNum, drawPeriod)
}

// calculate the turing pattern in general
func calculateTuringPatternBoard(patterns []TuringPattern, board GameBoard, stepNum, drawPeriod int) {
	// update the patterns[i].activator[][] and
	// 						inhibitor[i].inhibitor[][] for each turint scale
	var imglist []image.Image
	for step := 0; step < stepNum; step++ {
		fmt.Println(step)
		updateTuringScales(patterns, board)
		// use patterns[i].activator[][] and inhibitor[i].inhibitor[][] to
		//calculate variations variation=variation+abs(activator[x][y]-inhibitor[x][y])
		updateScalesVariation(patterns, board)
		updateBoardFromPatterns(patterns, board)
		normalizeBoard(board)
		// draw one image or add one image to gif at in the period of draw period
		if step%drawPeriod == 0 {
			image := drawGameBoard(board, step)
			imglist = append(imglist, image)
		}
	}
	// produce and concantnate the string for file name
	// the final format will be like "TuringPatternTotal20Step100SizeatEvery2steps.out.gif"
	imageName := fmt.Sprintf("TuringPatternTotal%dStep%dSizeatEvery%dsteps",
		stepNum, len(board), drawPeriod)
	process(imglist, imageName)
	fmt.Println("gif wrote")
}

// draw the current board state, save the result as .png
func drawGameBoard(board GameBoard, step int) image.Image {
	rows := len(board)
	cols := len(board[0])
	rowHeight := 1.0
	colWidth := 1.0
	//rowHeight := 2.0
	//colWidth := 2.0
	width := int(colWidth) * cols
	height := int(rowHeight) * rows
	// create a canvas
	boardCanvas := CreateNewCanvas(width, height)
	// draw
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// calculate the grey image
			// can also use color inage
			var greyScale = uint8((board[row][col] + 1.0) * 256.0)
			currColor := MakeColor(greyScale, greyScale, greyScale)
			boardCanvas.SetFillColor(currColor)
			drawCell(boardCanvas, row, col, rowHeight, colWidth)
		}
	}
	// save picture
	//name := fmt.Sprintf("TuringPatternStep %d.png", step)
	//boardCanvas.SaveToPNG(name)
	return boardCanvas.img
}

// draw cell
func drawCell(b Canvas, r, c int, h, w float64) {
	// get data
	x1 := float64(r) * h
	y1 := float64(c) * w
	x2 := float64(r+1) * h
	y2 := float64(c+1) * w
	// start to draw
	b.MoveTo(x1, y1)
	b.LineTo(x1, y2)
	b.LineTo(x2, y2)
	b.LineTo(x2, y1)
	b.LineTo(x1, y1)
	b.Fill()
}
