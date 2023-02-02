package main

import (
	"fmt"
	"go-cart/Soal1_2/compare"
	"go-cart/Soal1_2/pecahan"
	"os"

	"os/exec"
	"runtime"
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
	clearScreen()
	pilihan := 0

	fmt.Println("1. Menentukan Pecahan Uang")
	fmt.Println("2. Mengetahui string dapat diedit atau tidak")
	fmt.Println("3. Exit")
	fmt.Println("=========================================================")

	for {

		fmt.Print("Masukan Pilihan Anda : ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			pecahan.Pecahan()
			fmt.Println()
		case 2:
			compare.Compare()
			fmt.Println()
		case 3:
			fmt.Print("See You . . .")
			os.Exit(0)
		default:
			fmt.Println("Input yang anda masukan salah!")
			os.Exit(0)
		}

	}
}
