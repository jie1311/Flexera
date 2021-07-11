package main

import (
	"errors"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		err := errors.New("no filepath or application id provided")
		PrintError(err)
		return
	}

	copies, err := CalculateCopiesFromCsv(os.Args[1], os.Args[2])
	if err != nil {
		PrintError(err)
		return
	}

	PrintOutput(copies)
}
