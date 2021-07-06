package main

import (
	"fmt"
)

func filterKataKotor(s string) string {
	if s == "Anjing" {
		return "..."
	} else {
		return s
	}
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Halooo", filter(name))
}

func main() {
	sayHelloWithFilter("Denada", filterKataKotor)
	sayHelloWithFilter("Anjing", filterKataKotor)
}
