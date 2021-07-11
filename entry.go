package main

import (
	"fmt"
	"os"
)

func main() {
	copies, err := CalculateCopiesFromCsv(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println("Error(s) happened during processing, Please verify your input. More info:", err)
	} else {
		fmt.Println(copies, "copies are needed.")
	}
}
