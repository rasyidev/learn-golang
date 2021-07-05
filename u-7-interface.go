package main

import "log"

type Animal interface {
	Says() string
	getLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

// func main() {
// 	dog := Dog{
// 		Name:  "Samsul",
// 		Breed: "German Sephered",
// 	}

// 	PrintInfo(&dog)

// }

func (d *Dog) Says() string {
	return "Oug Oug"
}

func (d *Dog) getLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This Animal says", a.Says(), "and has", a.getLegs(), "legs.")
}
