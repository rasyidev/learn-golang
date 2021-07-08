package main

import "fmt"

type Orang struct {
	nama string
	umur int
}

type Siswa struct {
	tingkat string
	orang   Orang
}

func main() {

	orang1 := Orang{
		nama: "Andi Sukmo Raharjo",
		umur: 17,
	}

	orang2 := Orang{
		nama: "Budi Amin Sukmoharjo",
		umur: 20,
	}

	siswa1 := Siswa{
		tingkat: "SMA",
		orang:   orang1,
	}

	siswa2 := Siswa{
		tingkat: "SMA",
		orang:   orang2,
	}

	fmt.Println("Siswa 1", "nama:", siswa1.orang.nama, "Umur:", siswa1.orang.umur, "Tingkat:", siswa1.tingkat)
	fmt.Println("Siswa 2", "nama:", siswa2.orang.nama, "Umur:", siswa2.orang.umur, "Tingkat:", siswa2.tingkat)

}
