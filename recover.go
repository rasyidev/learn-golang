package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate2(s string) (bool, error) {
	if strings.TrimSpace(s) == "" {
		return false, errors.New("Input can not be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Input your name:")
	fmt.Scanln(&name)

	if valid, err := validate2(name); valid {
		fmt.Println("Hello, ", name)
	} else {
		panic(err.Error())
		fmt.Println("End")
	}
}
