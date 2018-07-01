//
// Copyright (c) 2018 by Aditya Malu adityamalu1@gmail.com. All Rights Reserved.
//

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//seeding the random number
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var board [4][4]int
	numbers := [2]int{2, 4}
	num1, num2 := boardinitnumbers(r1)
	row1, col1, row2, col2 := boardinitpositions(r1)

	board[row1][col1] = numbers[num1]
	board[row2][col2] = numbers[num2]
	PrintBoard(board)

	for {
		char := GetInput()
		switch {
		case (char == 'w'):
			fmt.Println("UP")
			up(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case char == 's':
			fmt.Println("DOWN")
			down(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case char == 'a':
			fmt.Println("LEFT")
			left(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case char == 'd':
			fmt.Println("RIGHT")
			right(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case char == 'q':
			os.Exit(0)
		default:
			fmt.Println("w for up, s for down, a for left, d for right, q to quit")
		}
	}

}

func boardinitpositions(r1 *rand.Rand) (int, int, int, int) {
	row1, col1 := r1.Intn(4), r1.Intn(4)
	row2, col2 := r1.Intn(4), r1.Intn(4)
	for (row1 == row2) && (col1 == col2) {
		row2, col2 = r1.Intn(4), r1.Intn(4)
	}
	return row1, col1, row2, col2
}

func boardinitnumbers(r1 *rand.Rand) (int, int) {
	return r1.Intn(2), r1.Intn(2)
}

// func PrintBoard(board [4][4]int) {
// 	var i, j int
// 	for i = 0; i < 4; i++ {
// 		for j = 0; j < 4; j++ {
// 			fmt.Print(" ", board[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }

func add(board *[4][4]int) {
	var i, j int
	for i = 0; i < 4; i++ {
		for j = 0; j < 3; j++ {
			if board[i][j] == board[i][j+1] {
				board[i][j] = board[i][j] + board[i][j+1]
				board[i][j+1] = 0
			}
		}
	}
}

func slide(board *[4][4]int) {
	var i, j int
	for i = 0; i < 4; i++ {
		var copyrow [4]int
		for j = 0; j < 4; j++ {
			if board[i][j] != 0 {
				temp := 0
				for copyrow[temp] != 0 {
					temp++
				}
				copyrow[temp] = board[i][j]
			}
		}
		board[i] = copyrow
	}
}

func clockwiserotate(board *[4][4]int) {
	copyboard := *board
	col := 3
	var i, j int
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			board[j][col] = copyboard[i][j]
		}
		col--
	}
}

func anticlockwiserotate(board *[4][4]int) {
	clockwiserotate(board)
	clockwiserotate(board)
	clockwiserotate(board)
}

func left(board *[4][4]int) {
	slide(board)
	add(board)
	slide(board)
}

func right(board *[4][4]int) {
	clockwiserotate(board)
	clockwiserotate(board)
	left(board)
	clockwiserotate(board)
	clockwiserotate(board)
}

func down(board *[4][4]int) {
	clockwiserotate(board)
	left(board)
	anticlockwiserotate(board)
}

func up(board *[4][4]int) {
	anticlockwiserotate(board)
	left(board)
	clockwiserotate(board)
}

func getblankpositions(board *[4][4]int, r1 *rand.Rand) (int, int) {
	var i, j, length int
	length = 0
	var blankpos [16][2]int
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			if board[i][j] == 0 {
				blankpos[length][0] = i
				blankpos[length][1] = j
				length++
			}
		}
	}
	blanktile := r1.Intn(length)
	tilerow := blankpos[blanktile][0]
	tilecol := blankpos[blanktile][1]
	return tilerow, tilecol
}

func fillablanktile(board *[4][4]int, r1 *rand.Rand, numbers [2]int) {
	blanktilerow, blanktilecol := getblankpositions(board, r1)
	board[blanktilerow][blanktilecol] = numbers[r1.Intn(2)]
}
