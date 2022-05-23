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
	pembulatan(4.32)
	var deret = nomor2.DeretBilangan{N: 30}

	// Nomor 2
	deret.Prima()
	deret.Ganjil()
	deret.Genap()
	deret.Fibonacci()

	// Nomor 3
	Length := nomor3.Kubus{Rusuk: 6.0}
	countKubus(&Length)
}

func pembulatan(num float64) string {
	s := ""
	if num == 0 {
		fmt.Println("error")
	} else {
		result := math.Round(num*10) / 10
		s := fmt.Sprintf("%.2f", result)
		fmt.Println(s)
		fmt.Println(reflect.TypeOf(s))
	}
	return s
}

func countKubus(hasil nomor3.Hitung) {
	k := fmt.Sprintf("%.2f", hasil.Keliling())
	fmt.Println("Keliling Balok: ", k)
	l := fmt.Sprintf("%.2f", hasil.Luas())
	fmt.Println("Luas Balok: ", l)
	v := fmt.Sprintf("%.2f", hasil.Volume())
	fmt.Println("Volume Balok: ", v)
}
