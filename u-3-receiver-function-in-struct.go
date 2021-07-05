package main

import (
	"fmt"
	"log"
)

type myStruct struct {
	FirstName string
}

// Asign function within a struct
func (m *myStruct) printFirstname() string {
	log.Println(m.FirstName)
	fmt.Println("Haloooooooo")
	return m.FirstName
}

// func main() {
// 	var myVar myStruct
// 	myVar.FirstName = "Millen"

// 	myVar2 := myStruct{
// 		FirstName: "Nikki",
// 	}

// 	log.Println("myVar value is:", myVar.printFirstname())
// 	log.Println("myVar2 value is:", myVar2.printFirstname())

// }
