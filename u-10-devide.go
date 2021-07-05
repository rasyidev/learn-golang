package main

import "log"

func Devide(a, b float32) (float32, error) {
	return a / b, nil
}

func main() {
	res, err := Devide(10.0, 1.0)

	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Result:", res)
}
