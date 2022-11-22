package main

import (
	"fmt"
)

func convert(Nilai int) {
	nilai := Nilai
	var huruf = [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	if nilai < 12 {
		fmt.Println(huruf[nilai])
	} else if nilai >= 12 && nilai <= 19 {
		fmt.Println(huruf[nilai-10] + " belas")
	} else if nilai >= 20 && nilai <= 99 {
		fmt.Println(huruf[(nilai-(nilai%10))/10] + " puluh " + huruf[nilai%10])
	} else if nilai >= 100 && nilai <= 199 {
		fmt.Print("seratus ")
		convert(nilai % 100)
	} else if nilai >= 200 && nilai <= 999 {
		fmt.Print(huruf[nilai/100] + " ratus ")
		convert(nilai % 100)
	} else if nilai >= 1000 && nilai <= 1999 {
		fmt.Print("seribu ")
		convert(nilai % 1000)
	} else if nilai >= 2000 && nilai <= 9999999 {
		fmt.Print(huruf[nilai/1000] + " ribu ")
		convert(nilai % 1000)
	} else if nilai >= 1000000 {
		fmt.Println("lebih dari 6 digit")
	}
}

func main() {
	convert(1)
}
