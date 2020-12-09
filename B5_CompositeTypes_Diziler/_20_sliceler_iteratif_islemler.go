package main

import "fmt"

func main() {
	/**
	 * Slice üzerinde iteratif işlemler
	 */
	var baskentler = []string  {"Ankara", "Delhi", "Tirana", "Bakü", "Roma"}
	var ulkeler = [] string {"Türkiye", "Hindistan", "Arnavutluk", "Azerbaycan", "İtalya"}
	var telefon_kodu = [] string {"+90", "+91", "+355", "+994", "+39"}

	i := 0
	for ; i<cap(baskentler); i++ {
		fmt.Println(ulkeler[i],"'nin başkenti ",baskentler[i],"'dir")
	}
	fmt.Println()
	for indis, deger := range telefon_kodu {
		fmt.Println(ulkeler[indis],"'nin telefon kodu: ", deger," dir")
	}
}
