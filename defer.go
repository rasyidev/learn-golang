package main

import "fmt"

func logging() {
	fmt.Println("Selesai menjalankan fungsi")
}

func devide(number int) {
	// defer function executed when number parent function executed or error
	fmt.Println("Deviding in progress . . .")
	defer logging()
	result := 10 / number
	fmt.Println("Result:", result)
}

func main() {
	defer logging()
	devide(10)
}
