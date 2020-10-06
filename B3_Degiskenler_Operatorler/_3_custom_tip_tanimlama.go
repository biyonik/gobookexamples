package main

import "fmt"

/**
 * Custom Tip Tanımlama
 * type tipismi veritipi
 */

type Saat int
type Dakika int
type Kilo float64
type Sehir string
type Cevap bool

func main() {
	saat := Saat(14)
 	dakika := Dakika(30)
 	kilo := Kilo(45.4)
 	sehir := Sehir("Isparta")
 	cevap := Cevap(true)

 	fmt.Printf("Saat %v:%v\n", saat, dakika)
 	fmt.Println("Kilo: ", kilo)
 	fmt.Println("Şehir: ", sehir)
 	fmt.Println("Cevap: ", cevap)
}
