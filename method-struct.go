package main

import "fmt"

type orang struct {
	namaDepan    string
	namaBelakang string
}

// getNamaOrang := (o *orang) func()

func (o *orang) getNamaLengkap() string {
	return o.namaDepan + " " + o.namaBelakang
}

func main() {
	orang1 := orang{
		namaDepan:    "Taylor",
		namaBelakang: "Swift",
	}

	namaLengkapOrang1 := orang1.getNamaLengkap()

	fmt.Println(namaLengkapOrang1)
}
