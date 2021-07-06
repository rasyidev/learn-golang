package main

import "log"

func main() {
	// for i := 0; i < 10; i++ {
	// 	log.Println("Hello ", i)
	// }

	myCats := []string{"Millen", "Nikki", "Olene"}

	// Key could be int, string
	// Use _ if key is not necessary
	for key, val := range myCats {
		log.Println(key, val)
	}

}
