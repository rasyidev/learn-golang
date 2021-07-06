package main

import "fmt"

func Apapun(i int) interface{} {
	switch i {
	case 1:
		return 100
	case 2:
		return "Halooo"
	case 3:
		return true
	case 4:
		return map[string]string{"name": "myname", "address": "MyAddress"}
	default:
		return "Yippie"
	}
}

func main() {
	result := Apapun(2)
	fmt.Println(result)
}
