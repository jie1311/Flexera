package main

import "testing"

func TestAddComputer(t *testing.T) {
	user := User{
		ID:        "1001",
		Computers: make(map[string]Computer),
	}
	user.AddComputer("9001", "LAPTOP")
	_, hasComputer := user.Computers["9001"]

	if len(user.Computers) == 0 ||
		!hasComputer ||
		user.Computers["9001"].ID != "9001" ||
		user.Computers["9001"].Type != "LAPTOP" {
		t.Error("AddComputer was incorrect.")
	}
}

func TestCountComputer(t *testing.T) {
	user := User{
		ID:        "1001",
		Computers: make(map[string]Computer),
	}
	user.Computers["9001"] = Computer{
		ID:   "9001",
		Type: "LAPTOP",
	}
	user.Computers["9002"] = Computer{
		ID:   "9002",
		Type: "DESKTOP",
	}

	desktops, laptops := user.CountComputer()

	if desktops != 1 || laptops != 1 {
		t.Error("CountComputer was incorrect.")
	}
}

func TestHasComputer(t *testing.T) {
	user := User{
		ID:        "1001",
		Computers: make(map[string]Computer),
	}
	user.Computers["9001"] = Computer{
		ID:   "9001",
		Type: "LAPTOP",
	}

	if !user.HasComputer("9001") {
		t.Error("HasComputer was incorrect.")
	}
}

func TestMinimalCopies(t *testing.T) {
	user := User{
		ID:        "1001",
		Computers: make(map[string]Computer),
	}
	user.Computers["9001"] = Computer{
		ID:   "9001",
		Type: "LAPTOP",
	}
	user.Computers["9002"] = Computer{
		ID:   "9002",
		Type: "DESKTOP",
	}
	user.Computers["9003"] = Computer{
		ID:   "9003",
		Type: "DESKTOP",
	}

	if user.MinimalCopies() != 2 {
		t.Error("MinimalCopies was incorrect.")
	}
}
