
# 02601TuringPattern
Final Project

Simulation of Turing Pattern using Golang.

How to Run:

1 necessary package:
The drawing module of go: 1) canvas.go; 2)gifhelper.go should be in the same folder
of the fp.go file


2 How to run and format of command line input
  1) set GOPATH
  2) cd to the folder containing the fp.go file
  3) build first--> go build
  4) run:
    The format of input is:   ./fp imageSize blurType stepNumber ScaleNumber drawPeriod
                             (./fp and 5 int)

  There are some input examples, which are suggested for play around and see the
  output pattern in image. During the running, the program will print step number
  which can save you some anxiety and thinking about the halting problem (which has
  been proved to be unsolvalble!)

  some worked examples:
      1)  ./fp
           1.txt    (simple mode with unstable pattern, very quick)
      2)  ./fp 200 0 800 3 20
           2.txt     (simple mode with final stable pattern, might need several minutes)
      3)  ./fp 300 2 50 3 5
           3.txt   (slower but with multiscale patterns, need several minutes )
      4)  ./fp 300 1 20 4 2    (rectangle mode with a little weird stable pattern
           4.txt                      need almost 10 minutes)
      5)  ./fp 400 2 200 3 10  (slower mode with nice large patterns. need 30+ minutes)
           5.txt



3)The meaning of the argument

      imageSize: An int. This number will specify dimensions of final output
                 .png image and .gif motion picture as imageSize by imageSize.
                 Considering the time substantial cost of drawing process the
                 requirement for obvious patter, a integer in range(256, 1024) is
                 suggested.

      blurType: An int, can only be 0 ,1 or 2. Each number represent a strategy
                to calculate a cell on board or a pixel in on final image.

                blurType = 0 ----> The fastest strategy. This strategy ill use
                                   row and col the cell is in to calculate the
                                  and update the value of the cell. Only 2*radius
                                  of cell will be considered
                blurType = 1 ----> A relative slower strategy. The value of a cell
                                  in next generation will be calculated from
                                  (2*radius)*(2*radius) cell around it. To be specific,
                                  if the cell we want to update is board[row][col],
                                  then, the cell in rectangle((row-radius, col-radius),
                                  (row+radius, col+radius)) will be used.
                blurType = 2 ----> The most reasonable and intuitive strategy.
                                   And unfortunately, is also much slower than the
                                   other two. For each cell board[row][col], the cell
                                   board[i][j] will be used if
                                   (i-row)^2 + (j-row)^2 < radius^2

      stepNumber: An int representing the how many time you want the board to be update
                  Usually 20 is good enough to yield a stable patterns

      scaleNum: An int with possible choice of 1, 2, 3, 4, 5.
                The Turing Pattern is simulated from a series of different dimensions
                Turing Scale(more specific explanation can be found in the report).
                The scaleNum specify how many different Turing Scale you want to use
                to simulate the Turing Pattern. There are at most 5 different scales
                in the code(kind of hardcoded). The more scales you use, the more time
                will be needed to run the code, because during update of each cell, a
                more layer of turing scale need to be calculated.


      drawPeriod: An int representing the interval of drawing. For example, an input
                  of 20 means you want to draw one image or add one scene to the final .gif
                  every 20 steps
