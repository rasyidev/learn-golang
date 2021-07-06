package main

import "fmt"

func eat(food string) {
	if food == "Gulai" {
		panic("Eits, gaboleh makan gulai ya. Kolesterol!")
	}
	fmt.Println("Asyik, lagi makan ", food, "nih")
}

// panic function stops a program, defer function will remain executed
func main() {
	eat("Bakso")
	eat("Gulai")
}
