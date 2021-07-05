package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasCat    bool   `json:"has_cat"`
}

func main() {
	myJSON := `
[
	{
		"first_name": "Peter",
		"last_name": "Parker",
		"hair_color": "Black",
		"has_cat": true
	},
	{
		"first_name": "Taylor",
		"last_name": "Swift",
		"hair_color": "Blonde",
		"has_cat": false
	}
]`
	// Read JSON
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJSON), &unmarshalled)

	if err != nil {
		log.Println("error unmarshall", err)
	}

	log.Println("Unmarshalled:", unmarshalled)

	// Write JSON

	var mySlice []Person
	p1 := Person{
		FirstName: "Avril",
		LastName:  "Lavigne",
		HairColor: "Blonde",
		HasCat:    true,
	}

	mySlice = append(mySlice, p1)
	newJSON, err := json.MarshalIndent(mySlice, "", "  ")

	if err != nil {
		log.Println("Error marshalling json:", err)
	}

	log.Println(string(newJSON))
}
