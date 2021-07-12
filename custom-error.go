package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(s string) (bool, error) {
	if strings.TrimSpace(s) == "" {
		newError := errors.New("The input can not be empty")
		return false, newError
	} else {
		return true, nil
	}
}

func main() {
	var input2 string
	fmt.Scanln(&input2)
	res, err := validate(input2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your input is valid:")
		fmt.Println(res)
	}
}
