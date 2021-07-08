package main

import "fmt"

func filterMax(numbers []int, max int) (int, func() []int) {
	var result []int

	for _, v := range numbers {
		if v <= max {
			result = append(result, v)
		}

	}
	return len(result), func() []int {
		return result
	}
}

func main() {
	numbers := []int{1, 3, 2, 6, 5, 2, 7, 9, 4, 3}

	numbersFiltered, getValues := filterMax(numbers, 4)
	theValues := getValues()
	fmt.Println("Number's Filtered:", numbersFiltered, "Values:", getValues())
	fmt.Println("Number's Filtered:", numbersFiltered, "Values:", theValues)
}
