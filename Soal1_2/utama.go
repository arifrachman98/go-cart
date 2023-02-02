package main

import (
	"fmt"
	"os"

	"os/exec"
	"runtime"

	"project_takehome/compare"
	"project_takehome/pecahan"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	
}