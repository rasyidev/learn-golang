package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Print("input a number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err != nil {
		fmt.Println(input, "is not a number")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Yup, its a number:", number)
	}

}
