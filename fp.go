package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Cell struct {
	curr float64 // represent the current board
	tmp  float64
	old  float64 // represent the old board
	diff float64 // represent the different value of the curr and the old
}

// The game board is a 2D slice of Cell objects//
type GameBoard [][]Cell

// The game board is a 2D slice of float64//
//type GameBoard [][]float64

type TuringPattern struct {
	activator                [][]float64
	inhibitor                [][]float64
	variations               [][]float64
	activatorRX, activatorRY int
	inhibitorRX, inhibitorRY int
	variationSampleR         int
	stepSize                 float64
	num                      int
	tmpDestA                 [][]float64
	tmpDestB                 [][]float64
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
	twoDSlice := make([][]float64, num)
	for i := 0; i < num; i++ {
		oneDSlice := make([]float64, num)
		twoDSlice[i] = oneDSlice
	}
	return twoDSlice
}

// initialize different scale turing patterns
func initializePatterns(patterns []TuringPattern, num, blurType, blurSteps int) []TuringPattern {
	var pattern0 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern0.initializeTuringPattern(num, 100, 100,
		200, 200, 1, blurType, blurSteps, 0.5)

	var pattern1 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern1.initializeTuringPattern(num, 50, 50,
		100, 100, 1, blurType, blurSteps, 0.4)

	var pattern2 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern2.initializeTuringPattern(num, 10, 10,
		20, 20, 1, blurType, blurSteps, 0.3)

	var pattern3 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern3.initializeTuringPattern(num, 5, 5,
		10, 10, 1, blurType, blurSteps, 0.2)

	var pattern4 TuringPattern
	//      initializeTuringPattern(num, aRX, aRY,
	//iRX, iRY, vSR, blurType, blurSteps int, stepSize float64)
	pattern4.initializeTuringPattern(num, 2, 2,
		4, 4, 1, blurType, blurSteps, 0.1)

	numOfPatterns := 5
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
	for i := 0; i < boardSize; i++ {
		boardRow := make([]Cell, boardSize)
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
			board[i][j].curr = rand.Float64()*2 - 1
		}
	}
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

func main() {
	// set up many variables
	boardSize := 256 // the dimensions of baord
	num := boardSize
	levels := 1
	blurType := 0
	blurSteps := 5
	//maxCount := 10000
	//convergeThreshold := 0.01 //1

	blurSteps, levels = updateBlurStepsAndLevels(blurSteps, levels)

	// patterns is a slice of different turing patterns
	patterns := make([]TuringPattern, 0)
	patterns = initializePatterns(patterns, num, blurType, blurSteps)

	// create the grid, with each cell{curr, tmp,  diff, old}
	turingBoard := createGameBoard(boardSize)
	// fill the board with random float64 in range(-1,1)
	turingBoard = initializeBoard(turingBoard)

	//diffBoardSum := 1

	calculateTuringPatternBoard(patterns, turingBoard)
}

// calculate the turing pattern in general
func calculateTuringPatternBoard(patterns []TuringPattern, board GameBoard) {
	// update the patterns[i].activator[][] and
	// 						inhibitor[i].inhibitor[][] for each turint scale

	updateTuringScales(patterns, board)
	// use patterns[i].activator[][] and inhibitor[i].inhibitor[][] to
	//calculate variations variation=variation+abs(activator[x][y]-inhibitor[x][y])
	updateScalesVariation(patterns, board)
	fmt.Println("turing Scale0=", patterns[0].variations)

}

// use patterns[i].activator[][] and inhibitor[i].inhibitor[][] to
//calculate variations variation=variation+abs(activator[x][y]-inhibitor[x][y])
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

func updateTuringScales(patterns []TuringPattern, board GameBoard) {
	// splash different update mode:
	//0--> fast mode;
	//1--> accurate mode;
	//2-->gussian mode(slow)

	switch patterns[0].updateType {
	case 0:
		quickUpdateScales(patterns, board)
	case 1:
		quickUpdateScales(patterns, board)
		//!!!! NEED TO WORK ON THS MODE
		//accurateUpdateScales(patterns, board) //
	case 2:
		quickUpdateScales(patterns, board)
		//!!!! NEED TO WORK ON THS MODE
		//gussianUpdateScales(patterns, board)
	}
}

///////////////////////////////////////////////////////////////////////
// quick update mode
///////////////////////////////////////////////////////////////////////

func quickUpdateScales(patterns []TuringPattern, board GameBoard) {
	scaleNum := len(patterns)
	for i := 0; i < scaleNum; i++ {
		patterns[i].quickUpdateScale(board)
	}
}

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

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for i := row - actR; i < row+actR; i++ {
				currRow := (i + rows) % rows
				sum += board[currRow][col].curr
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

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for j := col - actR; j < col+actR; j++ {
				currCol := (j + cols) % cols
				sum += board[row][currCol].curr
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

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for i := row - actR; i < row+actR; i++ {
				currRow := (i + rows) % rows
				sum += board[currRow][col].curr
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

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sum := 0.0
			for j := col - actR; j < col+actR; j++ {
				currCol := (j + cols) % cols
				sum += board[row][currCol].curr
			}
			turingScale.activator[row][col] =
				(turingScale.activator[row][col] + sum/float64(actR*2+1)) / 2.0
		}
	}
}

// check cell[i][j] is still in board
func inField(board GameBoard, i, j int) bool {
	rows, cols := len(board), len(board[0])
	if (i >= 0) && (i < rows) && (j >= 0) && (j < cols) {
		return true
	}
	return false
}

// draw the current board state, save the result as .png
func drawGameBoard(board GameBoard) /*image.Image */ {
	rows, cols := len(board), len(board[0])
	rowHeight, colWidth := 6.0, 6.0
	width, height := int(colWidth)*cols, int(rowHeight)*rows
	boardCanvas := CreateNewCanvas(width, height)
	myRed := MakeColor(255, 0, 0)
	myBlue := MakeColor(0, 0, 255)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col].curr != 0.0 {
				// if cooperate, blue cell
				boardCanvas.SetFillColor(myBlue)
			} else if board[row][col].curr != 0.0 {
				// if defect, red cell
				boardCanvas.SetFillColor(myRed)
			}
			drawCell(boardCanvas, row, col, rowHeight, colWidth)
		}
	}
	boardCanvas.SaveToPNG("Prisoners.png")
	//return prison.img
}

func drawCell(b Canvas, r, c int, h, w float64) {
	x1, y1 := float64(r)*h, float64(c)*w
	x2, y2 := float64(r+1)*h, float64(c+1)*w
	b.MoveTo(x1, y1)
	b.LineTo(x1, y2)
	b.LineTo(x2, y2)
	b.LineTo(x2, y1)
	b.LineTo(x1, y1)
	b.Fill()
}
