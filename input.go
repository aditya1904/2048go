//
// Copyright (c) 2018 by Aditya Malu adityamalu1@gmail.com. All Rights Reserved.
//

package main

import (
	"unicode"

	"github.com/eiannone/keyboard"
)

func GetInput() rune {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	dir := unicode.ToLower(char)
	// fmt.Println(char, reflect.TypeOf(char))
	return dir
	// fmt.Printf("You pressed: %q\r\n", char)
}
