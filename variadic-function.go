package main

import "fmt"

// sumAll accepts all number on its parameter
func sumAll(numbers ...int) int {
	res := 0
	for _, number := range numbers {
		res += number
	}
	return res
}

func main() {
	var a = sumAll(2, 3, 1, 4, 14, 5, 2)
	b := sumAll(5, 2, 5, 23, 6)

	fmt.Println("Nilai a:", a) // Nilai a: 31
	fmt.Println("Nilai b:", b) // Nilai b: 41
}
