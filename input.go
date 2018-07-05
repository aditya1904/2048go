//
// Copyright (c) 2018 by Aditya Malu adityamalu1@gmail.com. All Rights Reserved.
//

package main

import (
	"os"
	"os/exec"
)

func GetInput() []byte {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var b []byte = make([]byte, 3)
	os.Stdin.Read(b)
	return b
}
