package main

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

func CalculateCopiesFromCsv(path string, applicationID string) (copies int, err error) {
	input, err := os.Open(path)
	if err != nil {
		return
	}
	defer input.Close()

	reader := csv.NewReader(input)

	firstRecord, err := reader.Read()
	if err != nil {
		return
	}

	err = checkFormat(firstRecord)
	if err != nil {
		return
	}

	users := make(map[string]User)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if record[2] != applicationID {
			continue
		}

		if hasUser(users, record[1]) {
			if users[record[1]].HasComputer(record[0]) {
				continue
			}
			copies -= users[record[1]].MinimalCopies()
			users[record[1]].AddComputer(record[0], strings.ToUpper(record[3]))
			copies += users[record[1]].MinimalCopies()
		} else {
			users[record[1]] = User{
				ID:        record[1],
				Computers: make(map[string]Computer),
			}
			users[record[1]].AddComputer(record[0], strings.ToUpper(record[3]))
			copies += users[record[1]].MinimalCopies()
		}
	}

	if err != io.EOF {
		return
	} else {
		err = nil
	}
	return
}

func checkFormat(firstRecord []string) (err error) {
	if len(firstRecord) < 5 ||
		firstRecord[0] != "ComputerID" ||
		firstRecord[1] != "UserID" ||
		firstRecord[2] != "ApplicationID" ||
		firstRecord[3] != "ComputerType" ||
		firstRecord[4] != "Comment" {
		err = errors.New("file not supported")
	}
	return
}

func hasUser(users map[string]User, id string) (exsit bool) {
	_, exsit = users[id]
	return
}
