package main

import "fmt"

func PrintError(err error) {
	fmt.Println("Error(s) happened during processing, Please verify your input. More info:", err)
}

func PrintOutput(copies int) {
	fmt.Println(copies, "copies are needed.")
}
