//
// Copyright (c) 2018 by Aditya Malu adityamalu1@gmail.com. All Rights Reserved.
//

package main

import (
	"strconv"

	tm "github.com/buger/goterm"
)

var colors = map[int]int{
	2:    tm.CYAN,
	4:    tm.RED,
	8:    tm.GREEN,
	16:   tm.YELLOW,
	32:   tm.RED,
	64:   tm.GREEN,
	128:  tm.MAGENTA,
	256:  tm.BLUE,
	512:  tm.WHITE,
	1024: tm.MAGENTA,
	2048: tm.YELLOW,
}

func PrintBoard(board [4][4]int) {
	tm.Clear() // Clear current screen

	var i, j int
	x, y := tm.GetXY(40|tm.PCT, 30|tm.PCT)

	message := "Welcome to 2048. You know the rules. Play."
	tm.MoveCursor(x, y)
	tm.Println((tm.Color(tm.Bold(message), tm.CYAN)))

	tm.Print("\n")
	tm.MoveCursorForward(x - 5)
	tm.Printf((tm.Color(tm.Bold("w for up, s for down, a for left, d for right, q to quit"), tm.BLUE)))
	tm.Println("\n")
	tm.MoveCursorForward(x)
	tm.Print("+--------+--------+--------+--------+")

	for i = 0; i < 4; i++ {
		tm.Print("\n")
		tm.MoveCursorForward(x)
		tm.Print("|")
		tm.MoveCursorForward(2)

		for j = 0; j < 4; j++ {
			if board[i][j] == 0 {
				tm.Printf((tm.Color(tm.Bold("%4s"), colors[board[i][j]])), " ")
			} else {
				tm.Printf((tm.Color(tm.Bold("%4s"), colors[board[i][j]])), strconv.Itoa(board[i][j]))
			}
			tm.MoveCursorForward(2)
			tm.Print("|")
			tm.MoveCursorForward(2)
		}

		tm.Print("\n")
		tm.MoveCursorForward(x)
		tm.Print("+--------+--------+--------+--------+")
	}

	tm.Flush() // Call it every time at the end of rendering
	tm.MoveCursor(1, 1)
}
