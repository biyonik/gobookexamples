package main

import "fmt"

/**
 * Programın işleyişi sırasında birden fazla case adımının çalıştırılması
 * gereken durumlarla karşılaşılması durumunda case ifadesinden sonra
 * 'fallthrough' ifadesi kullanılır
 */
func main() {
	switch {
		case 2 != 8:
			fmt.Printf("İlk case adımı çalışıyor...\n")
			fallthrough
		case "Ali" == "Ahmet":
			fmt.Printf("İkinci case adımı çalışıyor...\n")
			fallthrough
		case 6*8 == 68:
			fmt.Printf("Üçüncü case adımı çalışıyor...\n")
		case 9-5 == 4:
			fmt.Printf("Dördüncü case adımı çalışıyor...\n")
	}
}
