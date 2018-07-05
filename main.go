//
// Copyright (c) 2018 by Aditya Malu adityamalu1@gmail.com. All Rights Reserved.
//

package main

import (
	"bytes"
	"fmt"
	"math/rand"
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

	up := []byte{27, 91, 65}
	down := []byte{27, 91, 66}
	left := []byte{27, 91, 68}
	right := []byte{27, 91, 67}

	for {
		b := GetInput()
		switch {
		case bytes.Equal(b, up):
			fmt.Println("UP")
			Up(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case bytes.Equal(b, down):
			fmt.Println("DOWN")
			Down(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case bytes.Equal(b, left):
			fmt.Println("LEFT")
			Left(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		case bytes.Equal(b, right):
			fmt.Println("RIGHT")
			Right(&board)
			fillablanktile(&board, r1, numbers)
			PrintBoard(board)
		default:
			fmt.Println("Press only Arrow Keys")
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

func Left(board *[4][4]int) {
	slide(board)
	add(board)
	slide(board)
}

func Right(board *[4][4]int) {
	clockwiserotate(board)
	clockwiserotate(board)
	Left(board)
	clockwiserotate(board)
	clockwiserotate(board)
}

func Down(board *[4][4]int) {
	clockwiserotate(board)
	Left(board)
	anticlockwiserotate(board)
}

func Up(board *[4][4]int) {
	anticlockwiserotate(board)
	Left(board)
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
