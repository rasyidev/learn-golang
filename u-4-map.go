package main

import "log"

func main() {

	var myMapOfString = make(map[int]string)
	myMapOfString[0] = "Zero"
	myMapOfString[1] = "One"
	myMapOfString[2] = "Two"

	type User struct {
		FirstName string
		LastName  string
	}

	var myMapOfUser = make(map[string]User)

	log.Println(myMapOfString)
	log.Println("zero:", myMapOfString[0], "one:", myMapOfString[1], "two:", myMapOfString[2])

	user1 := User{
		FirstName: "Avril",
		LastName:  "Lavigne",
	}

	user2 := User{
		FirstName: "Taylor",
		LastName:  "Swift",
	}

	myMapOfUser["al"] = user1
	myMapOfUser["ts"] = user2

	log.Println(myMapOfUser)
}
