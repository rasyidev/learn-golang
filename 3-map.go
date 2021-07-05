package main

import "fmt"

func main() {
	var Person = make(map[string]string) // Assign later (Empty Map)
	Person["name"] = "Woody Woodpacker"
	Person["city"] = "California, USA"

	var AnotherPerson = map[string]string{ // Assign right away
		"name": "Donald Duck",
		"city": "Las Vegas",
	}

	var Car = map[string]string{} // Assign Empty Map

	fmt.Println(Person)
	fmt.Println(AnotherPerson)
	fmt.Println(Car)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println(Person["name"], Person["city"])
	fmt.Println("Length of a Person's items is: ", len(Person))
}
