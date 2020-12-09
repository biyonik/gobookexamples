package main

import "fmt"

func main() {
	ogrenci_no := [7] int {101, 102, 103, 104, 105, 106, 107}
	ogrenci_not := [7] int {86, 26, 52, 100, 15, 5, 25}
	ortalama := 0
	toplam := 0

	for _, val := range ogrenci_not {
		toplam += val
	}

	ortalama = toplam / len(ogrenci_not)
	fmt.Println("Sınıfın ortalaması = ", ortalama)

	for indis := range ogrenci_not {
		if ogrenci_not[indis] > ortalama {
			fmt.Println(ogrenci_no[indis], " numaralı öğrencinin notu (",ogrenci_not[indis],") ortalamanın üzerinde!")
		}
	}
}
