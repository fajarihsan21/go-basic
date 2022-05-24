package main

import (
	"Tugas2/nomor2"
	"Tugas2/nomor3"
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Nomor 1
	pembulatan(4.35)

	// Nomor 2
	var deret = nomor2.DeretBilangan{N: 30}
	deret.Prima()
	deret.Ganjil()
	deret.Genap()
	deret.Fibonacci()

	// Nomor 3 (Hitung Kubus)
	Length := nomor3.Kubus{Rusuk: 6.0}
	countKubus(&Length)
}

// Nomor 1
func pembulatan(num float64) string {
	s := ""
	if num == 0 {
		fmt.Println("Angka tidak boleh nol")
	} else {
		result := math.Round(num*10) / 10
		s := fmt.Sprintf("%.2f", result)
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	}
	return s
}

// Nomor 3
func countKubus(hasil nomor3.Hitung) {
	k := fmt.Sprintf("%.2f", hasil.Keliling())
	fmt.Println("Keliling Kubus: ", k)
	l := fmt.Sprintf("%.2f", hasil.Luas())
	fmt.Println("Luas Kubus: ", l)
	v := fmt.Sprintf("%.2f", hasil.Volume())
	fmt.Println("Volume Kubus: ", v)
}
