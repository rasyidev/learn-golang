package main

import "log"

// func main() {
// 	log.Println("Hello world!")
// 	myCatName := "Millenia"
// 	log.Println("Memloc of myCatName is at", &myCatName)
// 	myCatName = "Nikki"
// 	log.Println("Memloc of myCatName after change the value is at", &myCatName)
// 	log.Println("-----------------------------------------------------------------")
// 	myKittenName := "Olene"
// 	log.Println("My kitten name is", myKittenName, "\tMemloc at", &myKittenName)
// 	changeMyKittenName(&myKittenName, "Pote")
// 	log.Println("My kitten name is", myKittenName, "\tMemloc at", &myKittenName)

// }

func changeMyKittenName(s *string, newName string) {
	log.Println("Kitten name will be change to", s)
	*s = newName
}
