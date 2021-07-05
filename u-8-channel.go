package main

import (
	"fmt"

	"github.com/rasyidev/learn_go/helpers"
)

func getValue(intChan chan int) {
	value := helpers.RandomNumber(20)
	intChan <- value
}

func main() {
	// log.Println("Hello from the main file of go")
	// helpers.Hello()
	intChan := make(chan int)
	defer close(intChan)

	go getValue(intChan)

	value := <-intChan

	fmt.Println("Value", value)

	// a := helpers.RandomNumber(10)
	// fmt.Println(a)
}
