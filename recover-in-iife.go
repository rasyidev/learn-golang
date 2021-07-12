package main

import "fmt"

func main() {
	data := []string{"Semar", "Bagong", "Petruk"}

	for _, person := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", person, "with message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			if person == "Semar" {
				panic("Aaaa lariii, ada semaaar!")
			}
		}()
	}
}
