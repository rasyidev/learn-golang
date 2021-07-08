package main

import "fmt"

func main() {
	getMinMax := func(numbers []int) (int, int) {
		var min int
		var max int
		for k, v := range numbers {
			switch {
			case k == 0:
				max = v
				min = v
			case v > max:
				max = v
			case v < min:
				min = v
			}
		}
		return min, max
	}

	numbers := []int{1, 3, 2, 8, 5, 4}
	min, max := getMinMax(numbers)

	fmt.Println("Min:", min, "Max:", max)
}
