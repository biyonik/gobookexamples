package main

import "fmt"

func main() {
	/**
	 * Golang programlama dilinde sadece for döngüsü vardır
	 * for <sayaç başlangıç>; <şart>; <iterasyon kriteri> { // döngü gövdesi }
	*/
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	sayac := 5
	us := 3
	sonuc := 1
	for; us > 0; us-- {
		sonuc *= sayac
	}
	fmt.Printf("%d'in 3. kuvveti := %d\n", sayac, sonuc)

	sayi := 64
	for i := 2; i <= sayi; i++ {
		if sayi % i == 0 {
			fmt.Printf("%d, %d'ün bölenidir\n", i, sayi)
		}
	}

}
