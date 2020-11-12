package main

import "fmt"

func main() {
	/**
	 * Özel if ifadesi kullanımı
	 * If bloğu başlangıcında tanımlanan değişken sadece if bloğu içinde kullanılır
	 * Özel koşul ifadesi tanımlanırken if yanındaki ifadede '()' kullanılmaz.
	 * Kullanılırsa syntax hatası alır.
	 */
	if i := 4; i % 2 == 0 {
		fmt.Printf("%d sayısı çifttir.\n", i)
	}

	if a := 5; a%2 == 0 {
		fmt.Printf("%d sayısı çifttir.\n", a)
	} else {
		fmt.Printf("%d sayısı tektir.\n", a)
	}
}
