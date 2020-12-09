package main

import "fmt"

func main() {
	/**
	 * var map_adi map [anahtar_deger_tipi] deger_tipi
	 */
	var bosMap map [string]int
	if bosMap == nil {
		 fmt.Println("'bosMap' boş bir map'tir")
	}

	map1 := map [string] int  {
		"Bir": 1,
		"İki": 2,
		"Üç": 3,
	}

	fmt.Println("Bir etiketli değer: ", map1["Bir"])
	fmt.Println("İki etiketli değer: ", map1["İki"])
	fmt.Println("Üç etiketli değer: ", map1["Üç"])
}
