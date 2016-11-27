package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The data stored in a single cell of a field
type Cell struct {
	strategy string  //represents "C" or "D" corresponding to the type of prisoner in the cell
	score    float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The game board is a 2D slice of Cell objects
type GameBoard [][]Cell

func main() {
	fileName, b, stepNumber := readParameters()
	file := openAndReadFile(fileName)
	//read file and get rows and cols from first line
	scanner := bufio.NewScanner(file)
	var rows, cols int
	rows, cols = getRowsAndCols(scanner)
	//read other lines and store the initial stragegy on board
	var board = make(GameBoard, rows)
	initializeBoard(scanner, board, rows, cols)
	//drawPrisonBoard(board)
	boardEvolution(board, b, stepNumber, fileName)
}

// update the whole prison
func boardEvolution(board GameBoard, b float64, stepNumber int, fileName string) {
	//var imglist []image.Image
	for stepCount := 0; stepCount < stepNumber; stepCount++ {
		scoreUpdate(board, b)
		evolution(board)
		//imageOfCurrentPrison := drawPrisonBoard(prison)
		//imglist = append(imglist, imageOfCurrentPrison)
	}
	//## need to draw if required
	//imageName := fileName + strconv.Itoa(stepNumber)
	//process(imglist, "image")
	drawPrisonBoard(board)
}

// update score according to current strategy of each cell
func scoreUpdate(board GameBoard, b float64) {
	rows, cols := len(board), len(board[0])
	//fmt.Println("rows, cols=", rows, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			//fmt.Println("row, col=", row, col, prison[row][col])
			updateAccordingToRule(board, row, col, b)
		}
	}
}

// update score of each cell, according to their strategies
func updateAccordingToRule(board GameBoard, row, col int, b float64) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if (i == row) && (j == col) {
				continue
			} else if !inField(board, i, j) { // how to deal with the edge and corner cells
				if board[row][col].strategy == "C" {
					board[row][col].score += 0
				} else if board[row][col].strategy == "D" {
					board[row][col].score += 0
				}
			} else {
				if (board[row][col].strategy == "C") && (board[i][j].strategy == "C") {
					board[row][col].score++
				} else if (board[row][col].strategy == "C") && (board[i][j].strategy == "D") {
					board[row][col].score += 0
				} else if (board[row][col].strategy == "D") && (board[i][j].strategy == "C") {
					board[row][col].score += b
				} else if (board[row][col].strategy == "D") && (board[i][j].strategy == "D") {
					board[row][col].score += 0
				}
			}
		}
	}
}

// strategy evolve according to rules
func evolution(board GameBoard) {
	rows, cols := len(board), len(board[0])
	var tempBoard = make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		var boardRows = make([]Cell, cols)
		for j := 0; j < cols; j++ {
			boardRows[j] = board[i][j]
		}
		tempBoard[i] = boardRows
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			//fmt.Println("row, cow=", row, col)
			maxStrategy := tempBoard[row][col].strategy
			maxScore := tempBoard[row][col].score
			for i := row - 1; i < row+2; i++ {
				for j := col - 1; j < col+2; j++ {
					if (inField(board, i, j)) && (maxScore < tempBoard[i][j].score) {
						maxStrategy = tempBoard[i][j].strategy
						maxScore = tempBoard[i][j].score
					}
				}
			}
			board[row][col].strategy = maxStrategy

		}
	}
	// after update strategy, clear score
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col].score = 0
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
func drawPrisonBoard(board GameBoard) /*image.Image */ {
	rows, cols := len(board), len(board[0])
	rowHeight, colWidth := 6.0, 6.0
	width, height := int(colWidth)*cols, int(rowHeight)*rows
	boardCanvas := CreateNewCanvas(width, height)
	myRed := MakeColor(255, 0, 0)
	myBlue := MakeColor(0, 0, 255)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col].strategy == "C" {
				// if cooperate, blue cell
				boardCanvas.SetFillColor(myBlue)
			} else if board[row][col].strategy == "D" {
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

// read files, return the each line as a string?
func openAndReadFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: something went wrong with the file")
	}
	return file
}

// read other lines and store the initial stragegy on board
func initializeBoard(scanner *bufio.Scanner, board GameBoard, rows, cols int) {
	count := 0
	for scanner.Scan() {
		var boardRows = make([]Cell, cols)
		for i := 0; i < rows; i++ {
			boardRows[i].strategy = string(scanner.Text()[i])
		}
		board[count] = boardRows
		count++
	}
	//fmt.Println("board[0], board[98]=", board[0], board[98])
}

// getRowsAndCols() from the first line if test
func getRowsAndCols(scanner *bufio.Scanner) (int, int) {
	count := 0
	var frows, fcols float64
	var rows, cols int

	for scanner.Scan() {
		if count == 0 {
			fmt.Sscanf(scanner.Text(), "%f %f", &frows, &fcols)
			rows = int(frows)
			cols = int(fcols)
			//fmt.Println("rows, cols = ", frows, fcols)
			//fmt.Println("rows, cols = ", rows, cols)
			//fmt.Println(scanner.Text())
			break
		}
	}
	return rows, cols
}

// readParameters() get input fron terminal
func readParameters() (string, float64, int) {
	// get parameters
	parameters := os.Args
	if parameters == nil || len(parameters) != 4 {
		fmt.Println("Error: 3 parameters are wanted.")
		os.Exit(1)
	}
	// parameter convert
	fileName := parameters[1]
	b, err2 := strconv.ParseFloat(parameters[2], 64)
	stepNumber, err3 := strconv.Atoi(parameters[3])
	// see if there is any illegal input
	if err2 != nil || b < 0 {
		fmt.Println("Error: err parameter b.")
		os.Exit(1)
	}
	if err3 != nil || stepNumber < 0 {
		fmt.Println("Error: err parameter stepnumber.")
		os.Exit(1)
	}
	return fileName, b, stepNumber
}
