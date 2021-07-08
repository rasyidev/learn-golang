package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2

}
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jariJari()
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("=== Persegi ===")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())

	fmt.Println("")

	bangunDatar = lingkaran{10.0}
	fmt.Println("=== Lingkaran ===")
	fmt.Println("luas		:", bangunDatar.luas())
	fmt.Println("keliling	:", bangunDatar.keliling())
	// Karena jariJari() gak ada di interface, perlu di-casting pake .(lingkaran).jariJari()
	fmt.Println("jari-jari	:", bangunDatar.(lingkaran).jariJari())

}
