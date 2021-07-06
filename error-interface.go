package main

import (
	"errors"
	"fmt"
)

func devide2(value int, devider int) (float64, error) {
	if devider == 0 {
		return 0.0, errors.New("Pembagian dengan nol")
	} else {
		return float64(value / devider), nil
	}
}

func main() {
	result, err := devide2(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
