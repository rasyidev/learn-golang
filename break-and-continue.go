package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println("Iterasi Ke: ", i)
	}
	fmt.Println("")
	/*
		Iterasi Ke:  0
		Iterasi Ke:  1
		Iterasi Ke:  2
	*/

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Iterasi Ke: ", i)
	}
	/*
		Iterasi Ke:  0
		Iterasi Ke:  1
		Iterasi Ke:  3
		Iterasi Ke:  4
	*/
}
