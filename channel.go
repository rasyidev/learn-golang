package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan string)

	sayHello := func(name string) {
		data := fmt.Sprintf("Halo %s", name)
		message <- data
	}

	go sayHello("Taylor Swift")
	go sayHello("Avril Lavigne")
	go sayHello("Bruno Mars")

	message1 := <-message
	fmt.Println(message1)
	message2 := <-message
	fmt.Println(message2)
	message3 := <-message
	fmt.Println(message3)

}
