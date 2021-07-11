package main

type User struct {
	ID        string
	Computers map[string]Computer
}

type Computer struct {
	ID   string
	Type string
}

func (u User) AddComputer(id, cType string) {
	u.Computers[id] = Computer{
		ID:   id,
		Type: cType,
	}
}

func (u User) CountComputer() (desktops, laptops int) {
	for _, computer := range u.Computers {
		if computer.Type == "DESKTOP" {
			desktops++
		}
	}
	laptops = len(u.Computers) - desktops
	return
}

func (u User) HasComputer(id string) (exsit bool) {
	_, exsit = u.Computers[id]
	return
}

func (u User) MinimalCopies() int {
	desktops, laptops := u.CountComputer()
	if desktops >= laptops {
		return desktops
	} else {
		return (desktops + laptops + 1) / 2
	}
}
